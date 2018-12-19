package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	projectName := "go-mega"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

