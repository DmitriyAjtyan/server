package main

import (
	"log"
	"net/http"
	route "server/back/route"
)

func main() {
	//GetConfiguration()
	route.RouterFunc()
	err := http.ListenAndServe(":8080", nil)
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
