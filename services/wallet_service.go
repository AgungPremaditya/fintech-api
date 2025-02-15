package services

import (
	wallet_dtos "ledger-system/dtos/wallet"
	"ledger-system/models"
	"ledger-system/repositories"
	"log"

	"github.com/shopspring/decimal"
)

type WalletService struct {
	walletRepo *repositories.WalletRepository
	userRepo   *repositories.UserRepository
	mapper     *wallet_dtos.Mapper
}

func NewWalletService(
	repo *repositories.WalletRepository,
	userRepo *repositories.UserRepository,
) *WalletService {
	return &WalletService{
		walletRepo: repo,
		userRepo:   userRepo,
		mapper:     wallet_dtos.NewWalletMapper(userRepo),
	}
}

func (s *WalletService) GetWalletsService() ([]wallet_dtos.WalletIndexDTO, error) {
	var wallets []models.Wallet

	wallets, err := s.walletRepo.GetWallets()
	if err != nil {
		log.Println("Error getting wallets:", err)
		return nil, err
	}

	return s.mapper.ToWalletResponseList(wallets), nil
}

func (s *WalletService) GetWalletService(id string) (wallet_dtos.WalletDetailDTO, error) {
	var wallet models.Wallet

	wallet, err := s.walletRepo.GetWallet(id)
	if err != nil {
		log.Println("Error getting wallet:", err)
		return wallet_dtos.WalletDetailDTO{}, err
	}

	balance, err := s.walletRepo.GetWalletBalance(id)
	if err != nil {
		log.Println("Error getting wallet balance:", err)
		return wallet_dtos.WalletDetailDTO{}, err
	}

	result := s.mapper.ToWalletDetailResponse(wallet, &balance)

	return result, nil
}

func (s *WalletService) CreateWalletService(wallet wallet_dtos.CreateWalletDTO) (*wallet_dtos.WalletDetailDTO, error) {
	newWallet, err := s.mapper.ToWalletModel(wallet)
	if err != nil {
		log.Println("Error mapping wallet:", err)
		return nil, err
	}

	createdWallet, err := s.walletRepo.CreateWallet(&newWallet)
	if err != nil {
		log.Println("Error creating wallet:", err)
		return nil, err
	}

	result := s.mapper.ToWalletDetailResponse(*createdWallet, &decimal.Zero)

	return &result, nil
}
