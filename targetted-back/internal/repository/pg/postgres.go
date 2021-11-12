package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	//nolint: misspell
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository/pg/migrations"
)

const (
	usersTable = "user_tb"
)

// NewPostgresDB returns new sqlx driver for postgres DB.
func NewPostgresDB(cfg config.Postgres) (*sqlx.DB, error) {
	pgConf := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.HOST, cfg.PORT, cfg.UserName, cfg.DBName, cfg.PAS, cfg.SSLMode)

	logrus.Infof("pg configurations:", pgConf)
	db, err := sqlx.Open("postgres", pgConf)
	if err != nil {
		return nil, fmt.Errorf("sqlx open connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping pg DB: %w", err)
	}

	err = migrations.RunPgMigrations()
	if err != nil {
		return nil, fmt.Errorf("run pg migrations: %w", err)
	}

	return db, nil
}
