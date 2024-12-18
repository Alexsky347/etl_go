package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
	"time"
)

const (
	ProductionEnv = "production"

	DatabaseTimeout    = 5 * time.Second
	ProductCachingTime = 1 * time.Minute
)

type Schema struct {
	Environment      string `env:"environment"`
	HTTPPort         int    `env:"http_port"`
	AuthSecret       string `env:"auth_secret"`
	LogLevel         string `env:"log_level"`
	Dsn              string `env:"database_url"`
	MessagingServer  string `env:"messaging_server"`
	MessagingTopic   string `env:"messaging_topic"`
	MessagingGroupId string `env:"messaging_group_id"`
	CSVPath          string `env:"csv_file_path"`
	ConfigPath       string `env:"config_file_path"`
	Separator        string `env:"separator"`
	OutputModeBatch  bool   `env:"output_mode_batch"`
}

var (
	cfg Schema
)

func LoadConfig() (*Schema, error) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(currentDir, "app-config.yml"))
	if err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Schema {
	return &cfg
}
