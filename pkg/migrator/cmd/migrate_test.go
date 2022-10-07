package cmd

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestNewMigrateCmd(t *testing.T) {
	options := &OptionsArg{
		FnGetDBCallback: getDBConnection,
		DdlFileStoreDir: "db/ddls",
		DmlFileStoreDir: "db/dmls",
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
