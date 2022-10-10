package ddl

import (
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
)

func createMigrationTemplate(name, version string) string {
	key := fmt.Sprintf("%s_%s", version, name)
	structName := migrator.SnakeCaseToCamelCase(fmt.Sprintf("%s_%s", name, version))
	return fmt.Sprintf(
		`package ddls

import (
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/ddl"
	"gorm.io/gorm"
)

func init() {ddl.Register(New%s)}
type %s struct {}
func New%s() migrator.IMigration {return &%s{}}
func (m *%s) Name() string {return "%s"}

// Up Called when "migrate ddl up" command
func (m *%s) Up(tx *gorm.DB) error{
	// Your gorm migration code goes here
	// Example: tx.Migrator().CreateTable(&models.User{})
	return nil
}

// Down Called when "migrate ddl down" command
func (m *%s) Down(tx *gorm.DB) error{
	// Your gorm migration code goes here
	// Example: tx.Migrator().DropTable(&models.User{})
	return nil
}
	`, structName, structName, structName, structName, structName, key, structName, structName,
	)
}

func createPkgTemplate(name string) string {
	return fmt.Sprintf(
		`package %s
`, name)
}
