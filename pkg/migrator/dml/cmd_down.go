package dml

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strings"
)

func ExecuteDown(args []string, fnGetDB migrator.FnDBConnection) {
	db := setUpMigration(fnGetDB)
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	num, needsConfirm, err := migrator.NumDownMigrationsFromArgs(viper.GetBool("dml.applyAll"), args)
	if err != nil {
		log.Fatal("error: ", err)
	}
	if needsConfirm {
		log.Println("Are you sure you want to apply all down migrations? [y/N]")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			return
		}
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" {
			log.Println("Applying all down migrations")
		} else {
			log.Fatal("Not applying all down migrations")
		}
	}

	if err := downCmd(db, num); err != nil {
		log.Fatal("error: ", err)
	}

	_ = versionCmd(db)
}

func downCmd(db *gorm.DB, limit int) error {
	counter := 0

	if len(seedersRun) <= 0 {
		return fmt.Errorf("no migration remain")
	}

	tx := db.Begin()
	for i := len(seedersRun) - 1; i >= 0; i-- {
		fnKey := seedersRun[i].Name
		v := seedersObjMap[fnKey]
		fmt.Println("Drop migration " + v.Name())
		err := v.Down(tx)
		if err != nil {
			tx.Rollback()
			log.Fatalf("%v", err)
		}

		migration := migration{Name: v.Name()}
		err = tx.Where("name = ?", v.Name()).Delete(&migration).Error
		if err != nil {
			tx.Rollback()
			log.Fatalf("%v", err)
		}
		fmt.Println("Migration " + v.Name() + " dropped successfully")

		if limit > 0 {
			limit = limit - 1
		}
		counter += 1

		if limit == 0 {
			break
		}
	}
	if res := tx.Commit(); res.Error != nil {
		log.Println("Migration Failed", res.Error)
		return res.Error
	}

	log.Printf("\n\n %d migration(s) dropped successfully \n\n", counter)
	return nil
}
