package pg

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	//nolint: misspell
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
)

const (
	usersTable = "user_tb"
)

type PgDb struct {
	*sqlx.DB
}

// NewPostgresDB returns new sqlx driver for postgres DB.
func NewPostgresDB(cfg config.Postgres) (*PgDb, error) {
	db, err := sqlx.Open(
		"pgx",
		cfg.DSN,
	)
	if err != nil {
		return nil, fmt.Errorf("connect to API postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping pg DB: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("run pg migrations: %w", err)
	}

	return &PgDb{db}, nil
}

var (
	errMigrationPath = errors.New("pg migration path not provided")
)

// runPgMigrations runs Postgres migrations
func (pg *PgDb) RunMigrations() error {
	cfg := config.Get() //!
	if cfg.Pg.MigrationsPath == "" {
		return errMigrationPath
	}

	m, err := migrate.New(
		cfg.Pg.MigrationsPath,
		cfg.Pg.DSN,
	)
	if err != nil {
		return fmt.Errorf("create pg migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate pg: %w", err)
	}
	return nil
}
