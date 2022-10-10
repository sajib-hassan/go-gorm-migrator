package cmd

import (
	"fmt"
	_ "github.com/sajib-hassan/go-gorm-migrator/example-app/db/ddls"
	_ "github.com/sajib-hassan/go-gorm-migrator/example-app/db/dmls"
	migratorCmd "github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/cmd"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func newMigrateCmd() *cobra.Command {
	options := &migratorCmd.OptionsArg{
		FnGetDBCallback: func() *gorm.DB {
			return getDBConnection()
		},
		// [OPTIONAL]
		//FileStoreDir: "db",
	}
	return migratorCmd.NewMigrateCmd(options)
}

func getDBConnection() *gorm.DB {
	dbHost := "localhost"
	dbPort := 5432
	dbName := "migrator_db"
	dbUser := "migrator_user"
	dbPassword := "migrator_password"
	dbSSLMode := "disable"
	dbTimeZone := "UTC"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, dbTimeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("postgres connect error: ", err)
	}
	return db
}
