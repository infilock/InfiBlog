package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/infilock/InfiBlog/config"

	"log"
)

func GetPsql(config *config.Psql) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.DSN())
	if err != nil {
		return nil, err //todo wrapcheck
	}

	if errPing := db.PingContext(context.Background()); errPing != nil {
		return nil, fmt.Errorf("error while pinging database: %w", errPing)
	}

	return db, nil
}

func MigrateUp(config *config.Psql) error {
	m, err := migrate.New(config.MigrationDir, config.URL())
	if err != nil {
		return fmt.Errorf("error while creating new migration: %w", err)
	}

	defer func(m *migrate.Migrate) {
		errSource, errDatabase := m.Close()
		if err != nil {
			log.Println("error close closes the source:", errSource)
			log.Println("error close closes the database:", errDatabase)
		}
	}(m)

	if errUp := m.Up(); errUp != nil && !errors.Is(errUp, migrate.ErrNoChange) {
		return fmt.Errorf("error while migrating up: %w", errUp)
	}

	return nil
}

func MigrateDown(config *config.Psql) error {
	m, err := migrate.New(config.MigrationDir, config.URL())
	if err != nil {
		return fmt.Errorf("error while creating new migration: %w", err)
	}

	defer func(m *migrate.Migrate) {
		errSource, errDatabase := m.Close()
		if err != nil {
			log.Println("error close closes the source:", errSource)
			log.Println("error close closes the database:", errDatabase)
		}
	}(m)

	if errDown := m.Down(); errDown != nil && !errors.Is(errDown, migrate.ErrNoChange) {
		return fmt.Errorf("error while migrating up: %w", errDown)
	}

	return nil
}
