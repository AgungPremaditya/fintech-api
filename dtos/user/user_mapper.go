package user_dtos

import (
	"ledger-system/models"
)

type Mapper struct{}

func NewUserMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) ToUserEmbedDTO(user models.User) UserEmbedDTO {
	return UserEmbedDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}
