package controllers

import (
	"encoding/json"
	"ledger-system/services"
	"net/http"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wallets, err := services.GetWalletsService()
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
