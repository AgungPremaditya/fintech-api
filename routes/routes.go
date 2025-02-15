package routes

import (
	"ledger-system/controllers"

	"github.com/gorilla/mux"
)

type Controllers struct {
	WalletController      *controllers.WalletController
	TransactionController *controllers.TransactionController
}

func SetupRoutes(c *Controllers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Wallet routes
	walletRouter := router.PathPrefix("/api/wallets").Subrouter()
	WalletRoutes(walletRouter, c.WalletController)

	// Transaction routes
	transactionRouter := router.PathPrefix("/api/transactions").Subrouter()
	TransactionRoutes(transactionRouter, c.TransactionController)

	return router
}

func WalletRoutes(router *mux.Router, c *controllers.WalletController) {
	router.HandleFunc("", c.GetWallets).Methods("GET")
	router.HandleFunc("", c.CreateWallet).Methods("POST")
	router.HandleFunc("/{id}", c.GetWallet).Methods("GET")
}

func TransactionRoutes(router *mux.Router, c *controllers.TransactionController) {
	router.HandleFunc("", c.CreateTransaction).Methods("POST")
	router.HandleFunc("/transfer", c.TransferTransaction).Methods("POST")
	router.HandleFunc("", c.GetTransactions).Methods("GET")
}
