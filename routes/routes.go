package routes

import (
	"ledger-system/controllers"

	"github.com/gorilla/mux"
)

type Controllers struct {
	WalletController *controllers.WalletController
}

func SetupRoutes(c *Controllers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Wallet routes
	walletRouter := router.PathPrefix("/wallets").Subrouter()
	WalletRoutes(walletRouter, c.WalletController)

	return router
}

func WalletRoutes(router *mux.Router, c *controllers.WalletController) {
	router.HandleFunc("", c.GetWallets).Methods("GET")
	router.HandleFunc("", c.CreateWallet).Methods("POST")
	router.HandleFunc("/{id}", c.GetWallet).Methods("GET")
}
