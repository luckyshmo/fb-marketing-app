package logger

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/sirupsen/logrus"
)

func Init(cfg config.Logging) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         cfg.SentryDSN,
		Environment: cfg.Environment,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel) //using debug lvl if we can't parse
		logrus.Warn("Using debug level logger")
	} else {
		logrus.SetLevel(lvl)
	}
	if cfg.Environment == "production" {
		var JSONF = new(logrus.JSONFormatter)
		JSONF.TimestampFormat = time.RFC3339
		logrus.SetFormatter(JSONF)
	}
}

func Info(args string) {
	sentry.CaptureMessage(args)
	logrus.Info(args)
}

func Error(args error) {
	sentry.CaptureException(args)
	logrus.Info(args)
}

func Warn(args error) {
	sentry.CaptureException(args)
	logrus.Warn(args)
}
