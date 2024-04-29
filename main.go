package main

import (
	"crud-go/config"
	"crud-go/controllers/homecontroller"
	"log"
	"net/http"
)

func main () {
	config.ConnectDB()

	//Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}