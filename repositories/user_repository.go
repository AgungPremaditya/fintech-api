package repositories

import (
	"fmt"
	"ledger-system/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUser(id string) (models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", id)

	if result.Error != nil {
		return models.User{}, fmt.Errorf("failed to fetch User: %v", result.Error)
	}

	return user, nil
}
