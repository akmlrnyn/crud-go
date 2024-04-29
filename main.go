package main

import (
	"crud-go/config"
	"crud-go/controllers/categorycontroller"
	"crud-go/controllers/homecontroller"
	"log"
	"net/http"
)

func main () {
	config.ConnectDB()

	//Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	//Categories Page
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)


	log.Println("Server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}