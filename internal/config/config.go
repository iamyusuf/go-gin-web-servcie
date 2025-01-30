package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type envConfigs struct {
	AppPort    int    `mapstructure:"APP_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int    `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
}

func (e *envConfigs) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		e.DbHost,
		e.DbUsername,
		e.DbPassword,
		e.DbName,
		e.DbPort,
	)
}

var EnvConfigs *envConfigs

func init() {
	EnvConfigs = loadEnv()
}

func loadEnv() (config *envConfigs) {
	viper.AddConfigPath(".")
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
