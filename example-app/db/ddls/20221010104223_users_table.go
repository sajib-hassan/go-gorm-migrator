package ddls

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/example-app/db/models"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/ddl"
	"gorm.io/gorm"
)

func init() { ddl.Register(NewUsersTable20221010104223) }

type UsersTable20221010104223 struct{}

func NewUsersTable20221010104223() migrator.IMigration { return &UsersTable20221010104223{} }
func (m *UsersTable20221010104223) Name() string       { return "20221010104223_users_table" }

// Up Called when "migrate ddl up" command
func (m *UsersTable20221010104223) Up(tx *gorm.DB) error {
	// Your gorm migration code goes here
	// Example: tx.Migrator().CreateTable(&models.User{})
	fmt.Println("up called")
	return tx.Migrator().CreateTable(&models.User{})
}

// Down Called when "migrate ddl down" command
func (m *UsersTable20221010104223) Down(tx *gorm.DB) error {
	// Your gorm migration code goes here
	// Example: tx.Migrator().DropTable(&models.User{})
	return tx.Migrator().DropTable(&models.User{})
}
