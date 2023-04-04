package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDatabase() error {
	fmt.Println("Migrating our database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("Cloud not create the postgres driver: %w", err)
	}

	Migrate, err := migrate.NewWithDatabaseInstance(
		"file:D:/Development/golang/RestApiV2/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
	}

	if err := Migrate.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("colud not run up migrations: %w", err)
		}

	}

	fmt.Println("migrated successfully migrated the db")

	return nil
}
