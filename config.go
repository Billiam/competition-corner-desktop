package main

import (
	"github.com/adrg/xdg"
  "github.com/spf13/viper"
  "fmt"
)

func LoadConfig() (error) {
  viper.SetDefault("width", 350)
  viper.SetDefault("height", 500)
  viper.SetDefault("x", 100)
  viper.SetDefault("y", 100)

	configFilePath, configErr := xdg.ConfigFile("competition-corner/config.json")
	if configErr != nil {
		return fmt.Errorf("could not resolve path for config file: %w", configErr)
	}

  viper.SetConfigFile(configFilePath)
  viper.SetConfigType("json")
  viper.ReadInConfig()

  return nil
}
