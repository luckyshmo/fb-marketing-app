package migrations

import (
	"errors"
	"fmt"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	errPgUrl         = errors.New("pg url not provided")
	errMigrationPath = errors.New("pg migration path not provided")
)

// runPgMigrations runs Postgres migrations
func RunPgMigrations() error { //? can be run from Makefile
	cfg := config.Get()
	if cfg.Pg.MigrationsPath == "" {
		return errMigrationPath
	}
	if !validUrl(cfg.Pg) {
		return errPgUrl
	}

	connectionString := "postgres://" + cfg.Pg.UserName + ":" + cfg.Pg.PAS +
		"@" + cfg.Pg.HOST + "/" + cfg.Pg.DBName + "?sslmode=" + cfg.Pg.SSLMode
	logrus.Info(cfg.Pg.MigrationsPath)
	logrus.Info(connectionString)
	m, err := migrate.New(
		cfg.Pg.MigrationsPath,
		connectionString,
	)
	if err != nil {
		return fmt.Errorf("create pg migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate pg: %w", err)
	}
	return nil
}

func validUrl(cfg config.Postgres) bool {
	return !(cfg.HOST == "" || cfg.PORT == "")
}
