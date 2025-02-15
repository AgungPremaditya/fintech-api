package services

import (
	transaction_dtos "ledger-system/dtos/transaction"
	"ledger-system/repositories"
	"log"
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

	log.Println("New transaction:", newTransaction)

	createdTransaction, err := s.transactionRepo.CreateTransaction(newTransaction)
	if err != nil {
		log.Println("Error creating transaction:", err)
		return nil, err
	}

	result := s.transactionMapper.ToTransactionDetailResponse(createdTransaction)

	return result, nil
}
