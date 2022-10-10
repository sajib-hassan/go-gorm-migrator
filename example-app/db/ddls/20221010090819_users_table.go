package ddls

import (
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/ddl"
	"gorm.io/gorm"
)

func init() { ddl.Register(NewUsersTable20221010090819) }

type UsersTable20221010090819 struct{}

func NewUsersTable20221010090819() migrator.IMigration { return &UsersTable20221010090819{} }
func (m *UsersTable20221010090819) Name() string       { return "20221010090819_users_table" }

func (m *UsersTable20221010090819) Up(tx *gorm.DB) error {
	// Your migration code goes here
	// Example: tx.Migrator().CreateTable(&models.User{})
	return nil
}

func (m *UsersTable20221010090819) Down(tx *gorm.DB) error {
	// Your migration code goes here
	// Example: tx.Migrator().DropTable(&models.User{})
	return nil
}
