package main

import (
	"log"
	"net/http"

	"go-crud-app/repository"
	"go-crud-app/router"
)

func main() {
	repository.InitDB()
	r := router.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
