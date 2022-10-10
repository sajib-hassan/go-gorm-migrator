package cmd

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestNewMigrateCmd(t *testing.T) {
	options := &OptionsArg{
		FnGetDBCallback: getDBConnection,
		FileStoreDir:    "db",
	}
	c := NewMigrateCmd(options)

	fmt.Printf("%v\n", c)
}

func getDBConnection() *gorm.DB {

	return &gorm.DB{
		Config:       nil,
		Error:        nil,
		RowsAffected: 0,
		Statement:    nil,
	}
}
