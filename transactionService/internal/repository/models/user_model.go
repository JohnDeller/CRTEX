package models

import (
	"database/sql"
	"time"
	"transactionService/pkg/proto/user_v1"
)

type UserModel struct {
	Id        string       `db:"id"`
	Balance   float64      `db:"balance"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func NewUserModel(user user_v1.User) UserModel {
	return UserModel{
		Id:      user.GetId(),
		Balance: user.Balance,
	}
}

func NewUserModelToDto(user UserModel) user_v1.User {
	return user_v1.User{
		Id:      user.Id,
		Balance: user.Balance,
	}
}
