package cmd

import (
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/ddl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func newDdlMigrateCmd(fnGetDB migrator.FnDBConnection, fileStoreDir string) *cobra.Command {

	// migrateCmd represents the migration command
	ddlCmd := &cobra.Command{
		Use:   "ddl",
		Short: "use DDL migration tool",
		Long:  `migration uses migrate tool under the hood supporting the same commands and an additional reset command`,
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
			ddl.ExecuteCreate(args)
		},
	}
	upCmd := &cobra.Command{
		Use:   "up",
		Short: "[N] Apply all or N up migrations",
		Run: func(cmd *cobra.Command, args []string) {
			ddl.ExecuteUp(args, fnGetDB)
		},
	}

	downCmd := &cobra.Command{
		Use:   "down",
		Short: "[N] [-all]    Apply all or N down migrations",
		Run: func(cmd *cobra.Command, args []string) {
			ddl.ExecuteDown(args, fnGetDB)
		},
	}

	dropCmd := &cobra.Command{
		Use: "drop",
		Short: `drop [-f]    Drop everything inside database
	Use -f to bypass confirmation`,
		Run: func(cmd *cobra.Command, args []string) {
			ddl.ExecuteDrop(args, fnGetDB)
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: `version      Print current migration version`,
		Run: func(cmd *cobra.Command, args []string) {
			ddl.ExecuteVersion(args, fnGetDB)
		},
	}

	ddlCmd.AddCommand(createCmd)
	ddlCmd.AddCommand(upCmd)
	ddlCmd.AddCommand(downCmd)
	ddlCmd.AddCommand(dropCmd)
	ddlCmd.AddCommand(versionCmd)

	fDir := "db/ddls"
	fileStoreDir = strings.TrimSpace(fileStoreDir)
	if len(fileStoreDir) > 0 {
		fDir = fileStoreDir
	}

	downCmd.Flags().Bool("all", false, "[-all] Apply all")
	_ = viper.BindPFlag("ddl.applyAll", downCmd.Flags().Lookup("all"))

	dropCmd.Flags().BoolP("force", "f", false, "Use -f to bypass confirmation")
	_ = viper.BindPFlag("ddl.forceDrop", dropCmd.Flags().Lookup("force"))

	createCmd.Flags().StringP("ext", "e", "go", "File extension")
	_ = viper.BindPFlag("ddl.extPtr", createCmd.Flags().Lookup("ext"))

	createCmd.Flags().StringP("dir", "", fDir,
		"Directory to place file in",
	)
	_ = viper.BindPFlag("ddl.dirPtr", createCmd.Flags().Lookup("dir"))

	createCmd.Flags().StringP("format", "f", "20060102150405",
		`The Go time format string to use. If the string "unix" or "unixNano" is specified, 
				then the seconds or nanoseconds since January 1, 1970 UTC respectively will be used. 
				Caution, due to the behavior of time.Time.Format(), invalid format strings will not error`,
	)
	_ = viper.BindPFlag("ddl.formatPtr", createCmd.Flags().Lookup("format"))

	createCmd.Flags().StringP("tz", "t", "UTC",
		`The timezone that will be used for generating timestamps (default: utc)`,
	)
	_ = viper.BindPFlag("ddl.timezoneName", createCmd.Flags().Lookup("tz"))

	createCmd.Flags().BoolP("seq", "q", false,
		"Use sequential numbers instead of timestamps (default: false)",
	)
	_ = viper.BindPFlag("ddl.seqNumber", createCmd.Flags().Lookup("seq"))

	createCmd.Flags().IntP("digits", "d", 6,
		"The number of digits to use in sequences (default: 6)",
	)
	_ = viper.BindPFlag("ddl.seqDigits", createCmd.Flags().Lookup("digits"))

	return ddlCmd
}
