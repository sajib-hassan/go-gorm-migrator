package cmd

import (
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/dml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func newDmlMigrateCmd(fnGetDB migrator.FnDBConnection, fileStoreDir string) *cobra.Command {

	// migrateCmd represents the migration command
	dmlCmd := &cobra.Command{
		Use:   "dml",
		Short: "use DB seed dml tool",
		Long:  `seeder uses migrate tool under the hood supporting the same commands and an additional reset command`,
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create [-ext E] [-dir D] [-seq] [-digits N] [-format] [-tz] NAME",
		Long: `create [-ext E] [-dir D] [-seq] [-digits N] [-format] [-tz] NAME
	   Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
	   Use -seq option to generate sequential up/down migrations with N digits.
	   Use -format option to specify a Go time format string. Note: migrations with the same time cause "duplicate migration version" error.
	   Use -tz option to specify the timezone that will be used when generating non-sequential migrations (defaults: UTC).
`,
		Run: func(cmd *cobra.Command, args []string) {
			dml.ExecuteCreate(args)
		},
	}
	upCmd := &cobra.Command{
		Use:   "up",
		Short: "[N] Apply all or N up migrations",
		Run: func(cmd *cobra.Command, args []string) {
			dml.ExecuteUp(args, fnGetDB)
		},
	}

	downCmd := &cobra.Command{
		Use:   "down",
		Short: "[N] [-all]    Apply all or N down migrations",
		Run: func(cmd *cobra.Command, args []string) {
			dml.ExecuteDown(args, fnGetDB)
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: `version      Print current migration version`,
		Run: func(cmd *cobra.Command, args []string) {
			dml.ExecuteVersion(args, fnGetDB)
		},
	}

	dmlCmd.AddCommand(createCmd)
	dmlCmd.AddCommand(upCmd)
	dmlCmd.AddCommand(downCmd)
	dmlCmd.AddCommand(versionCmd)

	fDir := "db"
	fileStoreDir = strings.TrimSpace(fileStoreDir)
	if len(fileStoreDir) > 0 {
		fDir = fileStoreDir
	}

	downCmd.Flags().Bool("all", false, "[-all] Apply all")
	_ = viper.BindPFlag("dml.applyAll", downCmd.Flags().Lookup("all"))

	createCmd.Flags().StringP("ext", "e", "go", "File extension")
	_ = viper.BindPFlag("dml.extPtr", createCmd.Flags().Lookup("ext"))

	createCmd.Flags().StringP("dir", "", fDir,
		"Directory to place file in",
	)
	_ = viper.BindPFlag("dml.dirPtr", createCmd.Flags().Lookup("dir"))

	createCmd.Flags().StringP("format", "f", "20060102150405",
		`The Go time format string to use. If the string "unix" or "unixNano" is specified,
				then the seconds or nanoseconds since January 1, 1970 UTC respectively will be used.
				Caution, due to the behavior of time.Time.Format(), invalid format strings will not error`,
	)
	_ = viper.BindPFlag("dml.formatPtr", createCmd.Flags().Lookup("format"))

	createCmd.Flags().StringP("tz", "t", "UTC",
		`The timezone that will be used for generating timestamps (default: utc)`,
	)
	_ = viper.BindPFlag("dml.timezoneName", createCmd.Flags().Lookup("tz"))

	createCmd.Flags().BoolP("seq", "q", false,
		"Use sequential numbers instead of timestamps (default: false)",
	)
	_ = viper.BindPFlag("dml.seqNumber", createCmd.Flags().Lookup("seq"))

	createCmd.Flags().IntP("digits", "d", 6,
		"The number of digits to use in sequences (default: 6)",
	)
	_ = viper.BindPFlag("dml.seqDigits", createCmd.Flags().Lookup("digits"))

	return dmlCmd
}
