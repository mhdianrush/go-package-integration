package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

var ENV *Config

func LoadConfig() {
	logger := logrus.New()
	logger.Println("Load Server Configuration")

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		panic(err)
	}
}
