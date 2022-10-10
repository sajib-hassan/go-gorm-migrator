package ddl

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func ExecuteUp(args []string, fnGetDB migrator.FnDBConnection) {
	db := setUpMigration(fnGetDB)
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	limit := -1
	if len(args) > 0 {
		n, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("error: can't read limit argument N")
		}
		limit = int(n)
	}

	if err := upCmd(db, limit); err != nil {
		log.Fatal("error: ", err)
	}

	_ = versionCmd(db)
}

func upCmd(db *gorm.DB, limit int) error {
	counter := 0

	tx := db.Begin()
	for _, fnKey := range migrationsFnMap {
		v := migrationsObjMap[fnKey]
		if isMigrationRunYet(v.Name(), migrationsRun) == false {
			fmt.Println("Run migration " + v.Name())
			err := v.Up(tx)
			if err != nil {
				tx.Rollback()
				log.Fatalf("%v", err)
			}

			migration := migration{Name: v.Name()}
			err = tx.Create(&migration).Error
			if err != nil {
				tx.Rollback()
				log.Fatalf("%v", err)
			}
			fmt.Println("Migration " + v.Name() + " run successfully")
			if limit > 0 {
				limit = limit - 1
			}
			counter += 1
		}
		if limit == 0 {
			break
		}
	}
	if res := tx.Commit(); res.Error != nil {
		log.Println("Migration Failed", res.Error)
		return res.Error
	}

	log.Printf("\n\n%d migration(s) run successfully \n\n", counter)

	return nil
}
