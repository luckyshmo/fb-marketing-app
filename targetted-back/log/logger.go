package logger

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
}

func Init(cfg config.Logging) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://c88f0b6955b74cf28f54389c3928b382@o918498.ingest.sentry.io/5862071",
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

	sentry.CaptureMessage("It works!")
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
