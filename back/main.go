package main

import (
	"fmt"
	"log"
	"net/http"
	config "server/back/config"
	postgres "server/back/postgres"
	redis "server/back/redis"
	route "server/back/route"
	"strconv"
)

func main() {
	config.GetServerConfig()
	postgres.Connect()
	log.Print("Successfuly connected to postgresql database")
	pool := redis.Connect()
	conn := pool.Get()
	defer conn.Close()
	err := redis.Ping(conn)
	if err != nil {
		fmt.Println(err)
	}
	log.Print("Successfuly connected to redis database")
	route.RouterFunc()
	log.Print("Server start at port: ", config.ConfigData.Port)
	err = http.ListenAndServe(fmt.Sprintf(":"+strconv.Itoa(config.ConfigData.Port)), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
