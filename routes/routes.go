package routes

import (
	"encoding/json"
	"ledger-system/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Health check route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": "Welcome to the Ledger System API",
		})
	}).Methods("GET")

	// Wallet routes
	router.HandleFunc("/wallets", controllers.GetWallets).Methods("GET")

	return router
}
