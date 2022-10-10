package dml

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"gorm.io/gorm"
	"log"
	"time"
)

var DefaultSeedersTable = "schema_seeders"

type migration struct {
	Name      string    `json:"name" gorm:"column:name;primarykey;size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"index:,type:brin"`
}

func (migration) TableName() string {
	return DefaultSeedersTable
}

var (
	seedersFnMap  []string
	seedersObjMap = make(map[string]migrator.IMigration)
	seedersRun    []migration
)

func getMigrations(db *gorm.DB) []migration {

	if db.Migrator().HasTable(DefaultSeedersTable) {
		err := db.Find(&seedersRun).Error
		if err != nil {
			fmt.Println("Can't loaded table migrations")
			panic(err)
		}

	} else {
		fmt.Println("Create table migrations")
		err := db.Migrator().CreateTable(&migration{})
		if err != nil {
			fmt.Println("Can't loaded table migrations")
			panic(err)
		}
	}

	return seedersRun
}

func isMigrationRunYet(name string, migrations []migration) bool {
	for _, v := range migrations {
		if v.Name == name {
			return true
		}
	}
	return false
}

func setUpMigration(fnGetDB migrator.FnDBConnection) *gorm.DB {
	db := fnGetDB()
	getMigrations(db)
	return db
}

func Register(m migrator.InitFn) {
	exists := false
	obj := m()
	fnName := obj.Name()
	for _, fn := range seedersFnMap {
		if fnName == fn {
			exists = true
		}
	}
	if exists {
		log.Fatalf("The dml %s has been already register", fnName)
	}
	seedersFnMap = append(seedersFnMap, fnName)
	seedersObjMap[fnName] = obj
}
