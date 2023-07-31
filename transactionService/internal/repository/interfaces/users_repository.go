package interfaces

import "transactionService/internal/repository/models"

type UsersRepository interface {
	Create(user models.UserModel) (string, error)
	Update(user models.UserModel) error
	UpdateBalance(userId string, userBalance float64) error
	GetById(userId string) (models.UserModel, error)
}
