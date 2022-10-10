package ddl

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strings"
)

func ExecuteDrop(_ []string, fnGetDB migrator.FnDBConnection) {
	if migrator.IsProductionEnv() {
		panic("NOT ALLOWED TO DROP IN PROD")
	}

	db := setUpMigration(fnGetDB)
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	forceDrop := viper.GetBool("ddl.forceDrop")
	if !forceDrop {
		log.Println("Are you sure you want to drop the entire database schema? [y/N]")
		var response string
		_, _ = fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" {
			log.Println("Dropping the entire database schema")
		} else {
			log.Fatal("Aborted dropping the entire database schema")
		}
	}

	if err := dropCmd(db); err != nil {
		log.Fatal("error: ", err)
	}
}

func dropCmd(db *gorm.DB) error {
	// select all tables in current schema
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema=(SELECT current_schema()) AND table_type='BASE TABLE'`
	tables, err := db.Raw(query).Rows()
	if err != nil {
		return err
	}
	defer func() {
		if errClose := tables.Close(); errClose != nil {
			log.Println("could not close table_schema")
		}
	}()

	// delete one table after another
	tableNames := make([]string, 0)
	for tables.Next() {
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			return err
		}
		if len(tableName) > 0 {
			tableNames = append(tableNames, tableName)
		}
	}
	if err := tables.Err(); err != nil {
		return err
	}

	if len(tableNames) > 0 {
		// delete one by one ...
		tx := db.Begin()
		for _, t := range tableNames {
			query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", t)
			if err := tx.Exec(query).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		if res := tx.Commit(); res.Error != nil {
			log.Println("Migration Failed", res.Error)
			return res.Error
		}

		log.Printf("\n\nDrop run successfully \n\n")
	}
	return nil
}
