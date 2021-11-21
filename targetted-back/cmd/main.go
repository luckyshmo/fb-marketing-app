package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/handler"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository/pg"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service"
	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/server"

	"github.com/sirupsen/logrus"
)

// @title Facebook ads manager
// @version 1.0
// @description Facebook ads manager

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
func run() (err error) {
	// config
	time.Sleep(time.Second)
	cfg := config.Get()
	// logger configuration
	logger.Init(cfg.Logging)
	defer sentry.Flush(2 * time.Second)

	//Init DB
	var repo *repository.Repository
	if cfg.FakeDB {
		repo = repository.NewInMemoryRepository()
	} else {
		db, err := pg.NewPostgresDB(cfg.Pg)
		if err != nil {
			return fmt.Errorf("failed to initialize db: %w", err)
		}
		repo = repository.NewSqlxRepository(db)
		defer func() {
			err = db.Close()
		}()
	}

	//Init main components
	services := service.NewService(repo, cfg)
	handlers := handler.NewHandler(services)

	//starting server
	srv := new(server.Backend)
	go func() {
		if err := srv.Run(cfg.AppPort, handlers.InitRoutes()); err != nil {
			logger.Error(fmt.Errorf("running http server: %w", err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	services.FacebookAPI.PageManager.CheckPageLimitTicker(ctx)

	logger.Info("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("App Shutting Down")

	return nil
}
