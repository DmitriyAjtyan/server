package main

import (
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", route.RouterHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
