package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string
	Postgres PostgresConfig
	HTTP     HTTPConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

type HTTPConfig struct {
	Host string
	Port string
}

func LoadConfig(name, fType, path string) (*Config, error) {
	viper.SetConfigName(name)
	viper.SetConfigType(fType)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Println("Config loaded successfully")
	return &cfg, nil
}
