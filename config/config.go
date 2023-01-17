package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER            string        `mapstructure:"DB_DRIVER"`
	DB_SOURCE            string        `mapstructure:"DB_SOURCE"`
	HTTPS_ADDRESS_CLIENT string        `mapstructure:"HTTPS_ADDRESS_CLIENT"`
	HTTP_ADDRESS_CLIENT  string        `mapstructure:"HTTP_ADDRESS_CLIENT"`
	SMTP_ADDRESS         string        `mapstructure:"SMTP_ADDRESS"`
	TokenSymetricKey     string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	TokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshDuration      time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfigClient(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
