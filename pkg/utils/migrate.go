package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func PostgresMigrate(db *sql.DB) error {
	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("with instance SQL error: %v\n", err)
	}

	src, err := (&file.File{}).Open("file://pkg/migrations")
	if err != nil {
		return fmt.Errorf("open migrations file error: %v\n", err)
	}

	m, err := migrate.NewWithInstance("file", src, "photo", instance)
	if err != nil {
		return fmt.Errorf("new with migrate instance error: %v\n", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("up migrate error: %v\n", err)
	}

	return nil
}
