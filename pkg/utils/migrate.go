package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func PostgresMigrate(db *sql.DB) error {
	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("%w: %v", consts.ErrMigrateInstance, err)
	}

	src, err := (&file.File{}).Open("file://migrations")
	if err != nil {
		return fmt.Errorf("%w: %v", consts.ErrMigrateOpenFile, err)
	}

	m, err := migrate.NewWithInstance("file", src, "photo", instance)
	if err != nil {
		return fmt.Errorf("%w: %v", consts.ErrMigrateNewInstance, err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("%w: %v", consts.ErrMigrateUp, err)
	}

	return nil
}
