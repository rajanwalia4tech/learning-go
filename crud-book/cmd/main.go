package main

import (
	"log"
	"net/http"

	"crud-go/pkg/api"
	"crud-go/pkg/model"
)

func main() {
	err := model.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := api.SetupRoutes()

	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
