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
	"gorm.io/gorm"
	"magic.pathao.com/veritas/common-pkg/migrator"
	"magic.pathao.com/veritas/common-pkg/migrator/ddl"
)

func init() {ddl.Register(New%s)}
type %s struct {}
func New%s() migrator.IMigration {return &%s{}}
func (m *%s) Name() string {return "%s"}
		
func (m *%s) Up(tx *gorm.DB) error{
	// Your migration code goes here
	// Example: tx.Migrator().CreateTable(&authorize.User{})
	return nil
}
		
func (m *%s) Down(tx *gorm.DB) error{
	// Your migration code goes here
	// Example: tx.Migrator().DropTable(&authorize.User{})
	return nil
}
	`, structName, structName, structName, structName, structName, key, structName, structName,
	)
}
