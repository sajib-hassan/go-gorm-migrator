package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	r := newRootCmd()
	if err := r.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "example-app",
		Short: "A Example APP for GORM Migrator",
		Long:  `A Example APP for GORM Migrator`,
	}

	rootCmd.AddCommand(
		newMigrateCmd(),
	)

	return rootCmd
}
