package controllers

import (
	"encoding/json"
	"ledger-system/services"
	"net/http"

	"github.com/gorilla/mux"
)

type WalletController struct {
	walletService *services.WalletService
}

func NewWalletController(service *services.WalletService) *WalletController {
	return &WalletController{
		walletService: service,
	}
}

func (c *WalletController) GetWallets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wallets, err := c.walletService.GetWalletsService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Failed to get wallets",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": wallets,
	})
}

func (c *WalletController) GetWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	wallet, err := c.walletService.GetWalletService(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Failed to get wallet",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": wallet,
	})
}
