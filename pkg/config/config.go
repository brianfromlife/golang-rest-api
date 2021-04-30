package config

import (
	"log"

	"github.com/spf13/viper"
)

type Settings struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPass     string `mapstructure:"DB_PASS"`
	Env        string `mapstructure:"ENV"`
	JwtExpires string `mapstructure:"JWT_EXPIRES"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

func New() *Settings {

	var cfg Settings

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Panicln("No env file, using environment variables.", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Error trying to unmarshal configuration", err)
	}

	return &cfg
}
