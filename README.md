# Go Gorm Migrator

GormMigrator is a straightforward migration assistant for [Gorm][gorm]. Gorm already provides
helpful [migrate capabilities][gormmigrate]; it simply lacks sufficient support for migration rollback and schema
versioning.

It has two distinct features:

- Migrations for the Data Definition Language (DDL)
- Migrations for the Data Manipulation Language (DML)

Here, the migration processes of [Ruby on Rails][ror] or [Laravel][laravel] have had a major influence on GormMigrator.

## Supported databases

It supports any of the [databases Gorm supports][gormdatabases] like:

- PostgreSQL
- MySQL
- SQLite
- Microsoft SQL Server

## Integration

**We have to ensure below items before run the `migrate` command:**

* Must have two directories `db/ddls` and `db/dmls` (directories are automatically created when a migration created).
* Must add package `ddls` and `dmls` accordingly
* Must import below packages in `cmd/migrate.go` or somewhere to preload the migrations.

```text
	_ "github.com/sajib-hassan/go-gorm-migrator/example-app/db/ddls"
	_ "github.com/sajib-hassan/go-gorm-migrator/example-app/db/dmls"
```

### Options

This is the options struct for the command integration, in case you don't want the defaults:

```go

type FnDBConnection func () *gorm.DB

type OptionsArg struct {
// Through a callback function, the migrator will obtain a database connection. 
// Here, an application database connection may be used. Return a `*gorm.DB` DB connection.  
FnGetDBCallback migrator.FnDBConnection
// [OPTIONAL]
FileStoreDir string // Default is "[APP_ROOT]/db"
}
```

### Example:

```go
// File Path: example-app/cmd/migrate.go

package cmd

import (
	"fmt"
	_ "github.com/sajib-hassan/go-gorm-migrator/example-app/db/ddls"
	_ "github.com/sajib-hassan/go-gorm-migrator/example-app/db/dmls"
	migratorCmd "github.com/sajib-hassan/go-gorm-migrator/pkg/migrator/cmd"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func newMigrateCmd() *cobra.Command {
	options := &migratorCmd.OptionsArg{
		FnGetDBCallback: func() *gorm.DB {
			return getDBConnection()
		},
		// [OPTIONAL]
		//FileStoreDir: "db",
	}
	return migratorCmd.NewMigrateCmd(options)
}

func getDBConnection() *gorm.DB {
	dbHost := "localhost"
	dbPort := 5432
	dbName := "migrator_db"
	dbUser := "migrator_user"
	dbPassword := "migrator_password"
	dbSSLMode := "disable"
	dbTimeZone := "UTC"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, dbTimeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("postgres connect error: ", err)
	}
	return db
}
```

## Usage

### Migrate

Available `migrate` sub commands

#### Commands

```shell
$ build/[APP_NAME] migrate      
```

```text
migration uses migrate tool under the hood supporting the same commands and an additional reset command

Usage:
  [APP_NAME] migrate [command]

Available Commands:
  ddl        use DDL migration tool
  dml        use DB seed data tool

Flags:
  -h, --help   help for migrate

```

##### Database version managed "data definition language" (ddl)

```shell
$ build/[APP_NAME] migrate ddl     
```

```text
migration uses migrate tool under the hood supporting the same commands and an additional reset command

Usage:
  [APP_NAME] migrate ddl [command]

Available Commands:
  create      create [-ext E] [-dir D] [-seq] [-digits N] [-format] [-tz] NAME
  down        [N] [-all]    Apply all or N down migrations
  drop        drop [-f]    Drop everything inside database
        Use -f to bypass confirmation
  up          [N] Apply all or N up migrations
  version     version      Print current migration version

Flags:
  -h, --help   help for ddl
```

```shell
$ build/[APP_NAME] migrate ddl create users_table
2022/10/10 10:42:23 .../[APP_NAME]/db/ddls/20221010104223_users_table.go
```

```shell
$ build/[APP_NAME] migrate ddl up [N]
$ build/[APP_NAME] migrate ddl down [N] [--all]
$ build/[APP_NAME] migrate ddl drop -f
$ build/[APP_NAME] migrate ddl version 
```

#### Database version managed table data

```shell
$ build/[APP_NAME] migrate dml     
```

```text
seeder uses migrate tool under the hood supporting the same commands and an additional reset command

Usage:
  [APP_NAME] migrate dml [command]

Available Commands:
  create      create [-ext E] [-dir D] [-seq] [-digits N] [-format] [-tz] NAME
  down        [N] [-all]    Apply all or N down migrations
  up          [N] Apply all or N up migrations
  version     version      Print current migration version

Flags:
  -h, --help   help for data
```

```shell
$ build/[APP_NAME] migrate dml create add_users
2022/10/10 11:24:24 .../[APP_NAME]/db/dmls/20221010112424_add_users.go
```

```shell
$ build/[APP_NAME] migrate dml up [N]
$ build/[APP_NAME] migrate dml down [N] [--all]
$ build/[APP_NAME] migrate dml version 
```

[gorm]: http://gorm.io/

[gormmigrate]: https://gorm.io/docs/migration.html

[gormdatabases]: https://gorm.io/docs/connecting_to_the_database.html

[ror]: https://rubyonrails.org/

[laravel]: https://laravel.com/