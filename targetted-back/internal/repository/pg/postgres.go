package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	//nolint: misspell
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository/pg/migrations"
)

const (
	usersTable = "user_tb"
)

// NewPostgresDB returns new sqlx driver for postgres DB.
func NewPostgresDB(cfg config.Postgres) (*sqlx.DB, error) {
	connectionString := "postgres://" + cfg.UserName + ":" + cfg.PAS +
		"@" + cfg.HOST + ":" + cfg.PORT + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode

	db, err := sqlx.Open(
		"pgx",
		connectionString,
	)
	if err != nil {
		return nil, fmt.Errorf("connect to API postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping pg DB: %w", err)
	}

	err = migrations.RunPgMigrations(connectionString)
	if err != nil {
		return nil, fmt.Errorf("run pg migrations: %w", err)
	}

	return db, nil
}
