package main

import (
	"ledger-system/controllers"
	"ledger-system/db"
	"ledger-system/repositories"
	"ledger-system/routes"
	"ledger-system/services"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize db
	initDB := db.Init()

	if len(os.Args) >= 2 {
		// Run database commands
		db.RunCommands(initDB, os.Args[1:])
	} else {
		// Initialize repositories
		repositories := repositories.NewRepositories(initDB)

		// Initialize services
		services := services.NewServices(repositories)

		controllers := &routes.Controllers{
			WalletController: controllers.NewWalletController(services.WalletService),
		}

		// Initialize routes
		router := routes.SetupRoutes(controllers)
		log.Fatal(http.ListenAndServe(":8080", router))

	}
}
