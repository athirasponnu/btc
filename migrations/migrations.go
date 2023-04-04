package migrations

import (
	"backend-code-base-template/config"
	"backend-code-base-template/internal/consts"
	"backend-code-base-template/internal/repo/driver"

	_ "github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"github.com/sirupsen/logrus"
)

func Migration(operation string) {
	// init the env config
	cfg, err := config.LoadConfig(consts.AppName)
	if err != nil {
		panic(err)
	}
	// logrus init
	log := logrus.New()

	// database connection
	pgsqlDB, err := driver.ConnectDB(cfg.Db)
	if err != nil {
		log.Fatalf("unable to connect the database")
		return
	}
	driverDB, err := postgres.WithInstance(pgsqlDB, &postgres.Config{})
	// migration instance creation
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/scripts",
		consts.DatabaseType, driverDB)
	if err != nil {
		panic(err)
	}
	if operation == "up" {
		err = m.Up()
	}
	if operation == "down" {
		err = m.Down()
	}
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	log.Printf("migration %v completed...", operation)

}
