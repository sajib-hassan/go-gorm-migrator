package cmd

import (
	migratorCmd "github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/cmd"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func newMigrateCmd() *cobra.Command {
	options := &migratorCmd.OptionsArg{
		FnGetDBCallback: func() *gorm.DB {
			return getDBConnection()
		},
	}
	return migratorCmd.NewMigrateCmd(options)
}

func getDBConnection() *gorm.DB {
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
	//	dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, dbTimeZone)
	//
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN: dsn,
	//}), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal("postgres connect error: ", err)
	//}
	//return db

	return &gorm.DB{
		Config:       nil,
		Error:        nil,
		RowsAffected: 0,
		Statement:    nil,
	}
}
