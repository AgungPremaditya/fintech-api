package controllers

import (
	"encoding/json"
	wallet_dtos "ledger-system/dtos/wallet"
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

func (c *WalletController) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var dto wallet_dtos.CreateWalletDTO

	// Mapping request bdata to DTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Invalid request payload",
		})
		return
	}

	// Create wallet
	wallet, err := c.walletService.CreateWalletService(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":          http.StatusInternalServerError,
			"error_message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": wallet,
	})
}
