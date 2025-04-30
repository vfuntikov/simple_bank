package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrateDB - runs all migrations in the migrations
func MigrateDB(migrationURL string, dbSource string) error {
	log.Println("migrating our database")
	m, err := migrate.New(
		migrationURL, dbSource)
	if err != nil {
		return fmt.Errorf("cannot create new migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no change on migration")
		} else {
			return fmt.Errorf("failed to run migrate up: %w", err)
		}
	}

	return nil
}
