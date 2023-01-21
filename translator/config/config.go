package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER          string        `mapstructure:"DB_DRIVER"`
	DB_SOURCE          string        `mapstructure:"DB_SOURCE"`
	HTTP_ADDRESS_TRANS string        `mapstructure:"HTTP_ADDRESS_TRANS"`
	TokenSymetricKey   string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	TokenDuration      time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfigTranslator(path string) (con Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&con)
	return
}
