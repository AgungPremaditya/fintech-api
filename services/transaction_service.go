package services

import (
	general_dtos "ledger-system/dtos/general"
	transaction_dtos "ledger-system/dtos/transaction"
	"ledger-system/models"
	"ledger-system/repositories"
	"log"
	"math"
)

type TransactionService struct {
	transactionRepo   *repositories.TransactionRepository
	walletRepo        *repositories.WalletRepository
	transactionMapper transaction_dtos.Mapper
}

func NewTransactionService(
	repo *repositories.TransactionRepository,
	walletRepo *repositories.WalletRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: repo,
		walletRepo:      walletRepo,
	}
}

func (s *TransactionService) CreateTransactionService(transaction transaction_dtos.CreateTransactionDTO) (*transaction_dtos.DetailTransactionDTO, error) {

	// Get wallet
	wallet, err := s.walletRepo.GetWallet(transaction.WalletID)
	if err != nil {
		log.Println("Error getting transaction:", err)
		return nil, err
	}

	// Map DTO to transaction models
	newTransaction, err := s.transactionMapper.ToTransactionModel(&transaction, &wallet)
	if err != nil {
		log.Println("Error mapping wallet:", err)
		return nil, err
	}

	createdTransaction, err := s.transactionRepo.CreateTransaction(newTransaction)
	if err != nil {
		log.Println("Error creating transaction:", err)
		return nil, err
	}

	result := s.transactionMapper.ToTransactionDetailResponse(createdTransaction)

	return result, nil
}

func (s *TransactionService) GetTransactionHistory(walletID string, pagination general_dtos.PaginationRequest) (*transaction_dtos.TransactionPaginatedDTO, error) {
	var transactions []models.Transaction

	// Get transactions
	transactions, err := s.transactionRepo.GetTransactionsByWalletId(walletID, pagination)
	if err != nil {
		log.Println("Error getting transactions:", err)
		return nil, err
	}

	if len(transactions) == 0 {
		return &transaction_dtos.TransactionPaginatedDTO{}, nil
	}

	// Set Data
	data := s.transactionMapper.ToTransactionListResponse(transactions)

	// Set Meta
	totalCount, err := s.transactionRepo.GetTransaction(walletID)
	if err != nil {
		log.Println("Error getting total count:", err)
		return nil, err
	}

	meta := general_dtos.PaginationMeta{
		Total:    *totalCount,
		PerPage:  pagination.PerPage,
		Page:     pagination.Page,
		LastPage: int(math.Ceil(float64(*totalCount) / float64(pagination.PerPage))),
	}

	return s.transactionMapper.ToTransactionPaginatedResponse(&data, &meta), nil
}
