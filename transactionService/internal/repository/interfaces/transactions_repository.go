package interfaces

import "transactionService/internal/repository/models"

type TransactionsRepository interface {
	Create(transaction models.TransactionModel) (string, error)
	GetById(transactionId string) (models.TransactionModel, error)
}
