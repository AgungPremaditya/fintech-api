package controllers

import (
	"encoding/json"
	transaction_dtos "ledger-system/dtos/transaction"
	"ledger-system/services"
	"net/http"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: service,
	}
}

func (c *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto transaction_dtos.CreateTransactionDTO

	// Mapping request body to DTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Invalid request payload",
		})
		return
	}

	// Create transaction
	createdTransaction, err := c.transactionService.CreateTransactionService(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":          http.StatusInternalServerError,
			"error_message": err.Error(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": createdTransaction,
	})
}
