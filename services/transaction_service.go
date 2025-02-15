package services

import (
	general_dtos "ledger-system/dtos/general"
	ledgerentry_dtos "ledger-system/dtos/ledger_entry"
	transaction_dtos "ledger-system/dtos/transaction"
	"ledger-system/models"
	"ledger-system/repositories"
	"log"
	"math"

	"github.com/shopspring/decimal"
)

type TransactionService struct {
	transactionRepo   *repositories.TransactionRepository
	walletRepo        *repositories.WalletRepository
	ledgerEntryRepo   *repositories.LedgerEntryRepository
	transactionMapper transaction_dtos.Mapper
	ledgerEntryMapper ledgerentry_dtos.Mapper
}

func NewTransactionService(
	repo *repositories.TransactionRepository,
	walletRepo *repositories.WalletRepository,
	ledgerEntryRepo *repositories.LedgerEntryRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: repo,
		walletRepo:      walletRepo,
		ledgerEntryRepo: ledgerEntryRepo,
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
	totalCount, err := s.transactionRepo.CountTransaction(walletID)
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

func (s *TransactionService) TransferTransactionService(transaction transaction_dtos.TransferTransactionDTO) (*ledgerentry_dtos.DetailTransferTransactionDTO, error) {
	// Get sender wallet
	fromWallet, err := s.walletRepo.GetWallet(transaction.FromWalletID)
	if err != nil {
		log.Println("Error getting transaction:", err)
		return nil, err
	}

	// Get receiver wallet
	toWallet, err := s.walletRepo.GetWallet(transaction.ToWalletID)
	if err != nil {
		log.Println("Error getting transaction:", err)
		return nil, err
	}

	// Set debit ledger entry
	debitEntry := &models.LedgerEntry{
		Type:     string(models.Debit),
		Amount:   decimal.NewFromFloat(transaction.Amount),
		Balance:  fromWallet.Balance.Sub(decimal.NewFromFloat(transaction.Amount)),
		WalletID: fromWallet.ID,
	}

	// Set credit ledger entry
	creditEntry := &models.LedgerEntry{
		Type:     string(models.Credit),
		Amount:   decimal.NewFromFloat(transaction.Amount),
		Balance:  toWallet.Balance.Add(decimal.NewFromFloat(transaction.Amount)),
		WalletID: toWallet.ID,
	}

	// Process ledger entry
	_, _, err = s.ledgerEntryRepo.ProcessTransferLedgerEntry(debitEntry, creditEntry)
	if err != nil {
		log.Println("Error processing ledger entry:", err)
		return nil, err
	}

	result := s.ledgerEntryMapper.ToTransactionTransferDetailResponse(debitEntry, creditEntry)

	return result, nil
}
