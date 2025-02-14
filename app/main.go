package main

import (
	"ledger-system/repositories"
	"ledger-system/routes"
	"log"
	"net/http"
)

func main() {
	repositories.Init()
	router := routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
