package app

import (
	"github.com/spf13/viper"
	"go-rest-api/helper"
)

func GetConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)
	return config
}
