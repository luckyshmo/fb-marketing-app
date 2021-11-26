package logger

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/sirupsen/logrus"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
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

	formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	formatter.File = true
	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.WithFields(logrus.Fields{
		"file": "main.go",
	})
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
