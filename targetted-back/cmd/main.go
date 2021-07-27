package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/handler"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/repository/pg"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/service"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/server"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// @title Example API
// @version 1.0
// @description API Server Example for building go microservices

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	//Mat Ryer advice to handle all app errors
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

//main func
func run() error {
	// config
	time.Sleep(time.Second)
	cfg := config.Get()
	// logger configuration
	logger.Init(cfg.Logging)
	defer sentry.Flush(2 * time.Second)

	//Init DB
	db, err := pg.NewPostgresDB(pg.Config{ //? you can get db by config
		Host:     cfg.PgHOST,
		Port:     cfg.PgPORT,
		Username: cfg.PgUserName,
		DBName:   cfg.PgDBName,
		SSLMode:  cfg.PgSSLMode,
		Password: cfg.PgPAS,
	})
	if err != nil {
		return errors.Wrap(err, "failed to initialize db")
	}

	logger.Error(fmt.Errorf("TEST ERROR"))
	logger.Info("TEST INFO")
	logger.Warn(fmt.Errorf("TEST WARN"))

	//Init main components
	//Good Clean arch and dependency injection example
	repos := repository.NewSqlxRepository(db)
	services := service.NewService(repos, cfg.Facebook)
	handlers := handler.NewHandler(services)

	//starting server
	srv := new(server.Server) //TODO? server.Server should be *serviceName*.server
	go func() {
		if err := srv.Run(cfg.AppPort, handlers.InitRoutes()); err != nil {
			logger.Error(fmt.Errorf("running http server: %w", err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	services.Facebook.CheckPageLimitTicker(ctx)

	logger.Info("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("App Shutting Down")
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
