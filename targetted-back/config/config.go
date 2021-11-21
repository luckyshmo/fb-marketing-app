package config

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
)

// Config. Should be filled from Env. Use launch.json(vscode) on local machine
type Config struct {
	AppPort    string `envconfig:"APP_PORT"`
	CorsHeader string `envconfig:"CORS_HEADER"`
	FakeDB     bool   `envconfig:"FAKE_DB"`

	Pg       Postgres
	Logging  Logging
	Facebook Facebook
	Youkassa Youkassa
}

// Postgres contains configurations info for Postgres DB.
type Postgres struct {
	HOST           string `envconfig:"PG_HOST"`
	PORT           string `envconfig:"PG_PORT"`
	PAS            string `envconfig:"PG_PAS"`
	SSLMode        string `envconfig:"PG_SSLMODE"`
	MigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	UserName       string `envconfig:"PG_USERNAME"`
	DBName         string `envconfig:"PG_DBNAME"`
}

type Facebook struct {
	Token       string `envconfig:"FB_TOKEN"`
	APIVersion  string `envconfig:"FB_API_VERSION"`
	BusinessID  string `envconfig:"FB_BUSINESS_ID"`
	AdAccountID string `envconfig:"AD_ACCOUNT_ID"`
}

type Youkassa struct {
	Username string `envconfig:"Y_UN"`
	Password string `envconfig:"Y_PAS"`
	ApiUrl   string `envconfig:"Y_URL"`
}

type Logging struct {
	// Logger
	LogLevel    string `envconfig:"LOG_LEVEL"`
	Environment string `envconfig:"ENV"`

	// Sentry
	SentryDSN string `envconfig:"SENTRY_DSN"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			logrus.Fatal(err)
		}
		validate(config)
	})
	logrus.Info(config)
	return &config
}

func validate(cfg Config) {
	refConf := reflect.ValueOf(cfg)
	typeOfRefConf := refConf.Type()

	for i := 0; i < refConf.NumField(); i++ {
		if fmt.Sprint(refConf.Field(i).Interface()) == "" {
			logrus.Error(fmt.Sprintf("Config: %s value is empty!", typeOfRefConf.Field(i).Name))
		}
	}
}
