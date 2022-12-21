package conf

import (
	"github.com/caarlos0/env/v6"
)

// AppConfig presents app conf
type AppConfig struct {
	AppEnv    string `env:"APP_ENV" envDefault:"dev"`
	Port      string `env:"PORT" envDefault:"8000"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"text"`
	DBHost    string `env:"DB_HOST" envDefault:""`
	DBPort    string `env:"DB_PORT" envDefault:"5432"`
	DBUser    string `env:"DB_USER" envDefault:""`
	DBPass    string `env:"DB_PASS" envDefault:""`
	DBName    string `env:"DB_NAME" envDefault:""`
	EnableDB  string `env:"ENABLE_DB" envDefault:"true"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
