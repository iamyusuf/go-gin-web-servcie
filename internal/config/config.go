package config

import (
	"github.com/spf13/viper"
	"log"
)

type envConfigs struct {
	AppPort int `mapstructure:"APP_PORT"`
}

var EnvConfigs *envConfigs

func init() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
