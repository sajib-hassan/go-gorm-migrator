package cmd

import (
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/spf13/cobra"
)

type OptionsArg struct {
	FnGetDBCallback migrator.FnDBConnection
	DdlFileStoreDir string
	DmlFileStoreDir string
}

func NewMigrateCmd(options *OptionsArg) *cobra.Command {
	// seedCmd to seed database
	// migrateCmd represents the migration command
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "use migration tool",
		Long:  `migration uses migrate tool under the hood supporting the same commands and an additional reset command`,
	}

	migrateCmd.AddCommand(
		newDdlMigrateCmd(options.FnGetDBCallback, options.DdlFileStoreDir),
	)

	return migrateCmd
}
