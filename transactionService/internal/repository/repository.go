package repository

import (
	"github.com/jmoiron/sqlx"
	"transactionService/internal/repository/interfaces"
	"transactionService/internal/repository/postgres"
)

type Repository struct {
	interfaces.TransactionsRepository
	interfaces.UsersRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TransactionsRepository: postgres.NewTransactionPostgres(db),
		UsersRepository:        postgres.NewUserPostgres(db),
	}
}
