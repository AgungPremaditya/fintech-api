package db

import (
	"ledger-system/models"
	"strconv"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func seedUsers(db *gorm.DB) error {
	// Seed users
	users := []models.User{
		{
			Name:     "John Doe",
			Email:    "johnDoe@mail.com",
			Password: "Secret123*",
		},
		{
			Name:     "Jane Doe",
			Email:    "janeDoe@mail.com",
			Password: "Secret123*",
		},
		{
			Name:     "Jack Doe",
			Email:    "jackDoe@mail.com",
			Password: "Secret123*",
		},
		{
			Name:     "Jill Doe",
			Email:    "jillDoe@mail.com",
			Password: "Secret123*",
		},
		{
			Name:     "Joe Doe",
			Email:    "joeDoe@mail.com",
			Password: "Secret123*",
		},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func seedWallets(db *gorm.DB) error {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return result.Error
	}

	for i, user := range users {
		wallet := models.Wallet{
			User:    user,
			Name:    "0x113F05B1123D71d591010282eEA07f34574c811" + strconv.Itoa(i),
			Balance: decimal.NewFromFloat(0.000000000000000001),
		}
		result := db.Create(&wallet)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func Seed(db *gorm.DB) error {
	// Seed users
	if err := seedUsers(db); err != nil {
		return err
	}

	// Seed wallets
	if err := seedWallets(db); err != nil {
		return err
	}

	return nil
}
