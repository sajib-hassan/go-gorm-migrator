package ddl

import (
	"errors"
	"fmt"
	"github.com/sajib-hassan/go-gorm-migrator/pkg/migrator"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ExecuteCreate(args []string) {
	startTime := time.Now()

	extPtr := viper.GetString("ddl.extPtr")
	dirPtr := viper.GetString("ddl.dirPtr")
	formatPtr := viper.GetString("ddl.formatPtr")
	timezoneName := viper.GetString("ddl.timezoneName")
	seq := viper.GetBool("ddl.seq")
	seqDigits := viper.GetInt("ddl.seqDigits")

	if len(args) == 0 {
		log.Fatal("error: please specify name")
	}

	name := args[0]

	if extPtr == "" {
		log.Fatal("error: --ext or -e flag must be specified")
	}

	timezone, err := time.LoadLocation(timezoneName)
	if err != nil {
		log.Fatal("error: ", err)
	}

	if err := createCmd(dirPtr, startTime.In(timezone), formatPtr, name, extPtr, seq, seqDigits, true); err != nil {
		log.Fatal("error: ", err)
	}
}

// createCmd (meant to be called via a CLI command) creates a new migration
func createCmd(dir string, startTime time.Time, format string, name string, ext string, seq bool, seqDigits int, print bool) error {
	if seq && format != migrator.DefaultTimeFormat {
		return migrator.ErrIncompatibleSeqAndFormat
	}

	var version string
	var err error

	dir = filepath.Clean(dir)
	ext = "." + strings.TrimPrefix(ext, ".")

	if seq {
		matches, err := filepath.Glob(filepath.Join(dir, "*"+ext))

		if err != nil {
			return err
		}

		version, err = migrator.NextSeqVersion(matches, seqDigits)

		if err != nil {
			return err
		}
	} else {
		version, err = migrator.TimeVersion(startTime, format)

		if err != nil {
			return err
		}
	}

	versionGlob := filepath.Join(dir, version+"_*"+ext)
	matches, err := filepath.Glob(versionGlob)

	if err != nil {
		return err
	}

	if len(matches) > 0 {
		return fmt.Errorf("duplicate migration version: %s", version)
	}

	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	basename := "ddls"
	filename := filepath.Join(dir, basename+ext)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		createPkgFile(filename, basename)
	}

	basename = fmt.Sprintf("%s_%s%s", version, name, ext)
	filename = filepath.Join(dir, basename)

	if err = createFile(filename, name, version); err != nil {
		return err
	}

	if print {
		absPath, _ := filepath.Abs(filename)
		log.Println(absPath)
	}

	return nil
}

func createFile(filename string, name, version string) error {
	// create exclusive (fails if file already exists)
	// os.Create() specifies 0666 as the FileMode, so we're doing the same
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Can't close file")
			panic(err)
		}
	}(f)

	if err != nil {
		return err
	}

	_, err = f.WriteString(createMigrationTemplate(name, version))
	if err != nil {
		return err
	}
	return nil
}

func createPkgFile(filename, name string) error {
	// create exclusive (fails if file already exists)
	// os.Create() specifies 0666 as the FileMode, so we're doing the same
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Can't close pkg file")
			panic(err)
		}
	}(f)

	if err != nil {
		return err
	}

	_, err = f.WriteString(createPkgTemplate(name))
	if err != nil {
		return err
	}
	return nil
}
