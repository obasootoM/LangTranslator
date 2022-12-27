package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER        string    `mapstructure:"DB_DRIVER"`
	DB_SOURCE        string    `mapstructure:"DB_SOURCE"`
	HTTPS_ADDRESS    string    `mapstructure:"HTTPS_ADDRESS"`
	HTTP_ADDRESS     string    `mapstructure:"HTTP_ADDRESS"`
	TokenSymetricKey string    `mapstructure:"TOKEN_SYMETRIC_KEY"`
	TokenDuration    time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
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
