package dmls

import (
	"github.com/go-faker/faker/v4"
	"github.com/sajib-hassan/go-gorm-migrator/example-app/db/models"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/dml"
	"gorm.io/gorm"
)

func init() { dml.Register(NewAddUsers20221010112424) }

type AddUsers20221010112424 struct{}

func NewAddUsers20221010112424() migrator.IMigration { return &AddUsers20221010112424{} }
func (m *AddUsers20221010112424) Name() string       { return "20221010112424_add_users" }

func (m *AddUsers20221010112424) Up(tx *gorm.DB) error {
	// Your gorm dml code goes here
	// Example: err := tx.Save(&datum).Error
	for i := 0; i < 10; i++ {
		u := models.User{}
		err := faker.FakeData(&u)
		if err == nil {
			err = tx.Save(&u).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *AddUsers20221010112424) Down(tx *gorm.DB) error {
	// Your gorm dml code goes here
	// Example: tx.Unscoped().Where(&datum).Delete(&datum).Error
	err := tx.Unscoped().Where("1 = 1").Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
