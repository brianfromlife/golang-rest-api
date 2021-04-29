package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Settings struct {
	DbHost string
	DbPort int32
	DbName string
	DbUser string
	DbPass string
}

func New() *Settings {

	var cfg Settings

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No env file, using environment variables.", err)
	}

	cfg.DbHost = viper.GetString("RDS_HOSTNAME")
	cfg.DbPort = viper.GetInt32("RDS_PORT")
	cfg.DbName = viper.GetString("RDS_DB_NAME")
	cfg.DbUser = viper.GetString("RDS_USERNAME")
	cfg.DbPass = viper.GetString("RDS_PASSWORD")

	return &cfg
}
