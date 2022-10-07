package ddl

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"gorm.io/gorm"
	"log"
	"time"
)

var DefaultMigrationsTable = "schema_migrations"

type migration struct {
	Name      string    `json:"name" gorm:"column:name;primarykey;size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"index:,type:brin"`
}

func (migration) TableName() string {
	return DefaultMigrationsTable
}

var (
	migrationsFnMap  []string
	migrationsObjMap = make(map[string]migrator.IMigration)
	migrationsRun    []migration
)

func getMigrations(db *gorm.DB) []migration {

	if db.Migrator().HasTable(DefaultMigrationsTable) {
		err := db.Find(&migrationsRun).Error
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

	return migrationsRun
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
	for _, fn := range migrationsFnMap {
		if fnName == fn {
			exists = true
		}
	}
	if exists {
		log.Fatalf("The migration %s has been already register", fnName)
	}
	migrationsFnMap = append(migrationsFnMap, fnName)
	migrationsObjMap[fnName] = obj
}
