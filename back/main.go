package main

import (
	"fmt"
	"log"
	"net/http"
	config "server/back/config"
	route "server/back/route"
	"strconv"
)

func main() {
	config.GetServerConfig()
	route.RouterFunc()
	log.Print("Server start at port: ", config.ConfigData.Port)
	err := http.ListenAndServe(fmt.Sprintf(":"+strconv.Itoa(config.ConfigData.Port)), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// GetConfiguration is a main configuration function of project
/*
func GetConfiguration() {
	viper.AddConfigPath("/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s", err)
	}
}
*/
