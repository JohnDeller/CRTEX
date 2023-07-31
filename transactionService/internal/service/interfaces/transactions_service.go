package interfaces

import (
	"transactionService/pkg/proto/transaction_v1"
)

type TransactionsService interface {
	Create(transaction transaction_v1.Transaction) (string, error)
	GetById(transactionId string) (transaction_v1.Transaction, error)
}
