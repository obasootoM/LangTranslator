package config

import (
	"time"

	"github.com/spf13/viper"
)





type Config struct {
	DB_DRIVER                string        `mapstructure:"DB_DRIVER"`
	DB_SOURCE_TRANSLATOR     string        `mapstructure:"DB_SOURCE_TRANSLATOR"`
	HTTP_ADDRESS_TRANSLATOR  string        `mapstructure:"HTTP_ADDRESS_TRANSLATOR"`
	TokenSymetricKey         string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	TokenDuration            time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshDuration          time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	HTTPS_ADDRESS_TRANSLATOR string        `mapstructure:"HTTPS_ADDRESS_TRANSLATOR"`
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