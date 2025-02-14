package routes

import (
	"ledger-system/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Health check route
	router.HandleFunc("/", controllers.GetApiHealth).Methods("GET")

	// Wallet routes
	router.HandleFunc("/wallets", controllers.GetWallets).Methods("GET")
	router.HandleFunc("/wallets/{id}", controllers.GetWallet).Methods("GET")

	return router
}
