package interfaces

import (
	"transactionService/pkg/proto/user_v1"
)

type UsersService interface {
	Create(user user_v1.User) (string, error)
	Update(user user_v1.User) error
	UpdateBalance(userId string, userBalance float64) error
	GetById(userId string) (user_v1.User, error)
}
