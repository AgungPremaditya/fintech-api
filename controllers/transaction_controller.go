package controllers

import (
	"encoding/json"
	general_dtos "ledger-system/dtos/general"
	transaction_dtos "ledger-system/dtos/transaction"
	"ledger-system/services"
	"net/http"
	"strconv"
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

func (c *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	// Set Page
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	// Set page size
	perPageStr := r.URL.Query().Get("page_size")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = 10
	}

	// Get transactions
	walletID := r.URL.Query().Get("wallet_id")
	transactions, err := c.transactionService.GetTransactionHistory(walletID, general_dtos.PaginationRequest{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":          http.StatusInternalServerError,
			"error_message": err.Error(),
		})
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": transactions,
	})
}
