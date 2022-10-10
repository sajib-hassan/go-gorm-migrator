package dml

import (
	"errors"
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ExecuteVersion(_ []string, fnGetDB migrator.FnDBConnection) {
	db := setUpMigration(fnGetDB)
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	if err := versionCmd(db); err != nil {
		log.Fatal("error: ", err)
	}
}

func versionCmd(db *gorm.DB) error {
	lastOne := &migration{}
	db.Logger = logger.Discard
	defer func() {
		db.Logger = logger.Default
	}()

	res := db.Last(lastOne)

	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		log.Println("Can't loaded last record of table migrations")
		return res.Error
	}
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	fmt.Printf("\nVersion: %s \nRun at: %s \n", lastOne.Name, lastOne.CreatedAt)
	return nil
}
