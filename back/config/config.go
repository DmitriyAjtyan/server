package config

import (
	"log"

	viper "github.com/spf13/viper"
)

// ConfigData is a variable contain server configuration
var ConfigData ConfigurationDataStruct

// GetServerConfig is a function which get configuration of server from json file
func GetServerConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../back/.")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&ConfigData)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
