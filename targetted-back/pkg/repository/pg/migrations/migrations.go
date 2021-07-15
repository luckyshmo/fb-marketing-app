package migrations

import (
	"errors"
	"fmt"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"

	"github.com/golang-migrate/migrate/v4"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// runPgMigrations runs Postgres migrations
func RunPgMigrations() error { //? can be run from Makefile
	cfg := config.Get()
	if cfg.PgMigrationsPath == "" {
		logger.Error(fmt.Errorf("migration path not provided"))
		return nil
	}
	if cfg.PgHOST == "" || cfg.PgPORT == "" {
		return errors.New("no cfg.PgURL provided")
	}

	connectionString := "postgres://" + cfg.PgUserName + ":" + cfg.PgPAS + "@" + cfg.PgHOST + "/" + cfg.PgDBName + "?sslmode=" + cfg.PgSSLMode

	m, err := migrate.New(
		cfg.PgMigrationsPath,
		connectionString,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
