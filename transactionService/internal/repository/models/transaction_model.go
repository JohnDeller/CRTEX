package models

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
	"transactionService/pkg/proto/transaction_v1"
)

type TransactionModel struct {
	Id            string       `db:"id"`
	UserId        string       `db:"user_id"`
	OperationType string       `db:"operation_type"`
	Price         float64      `db:"price"`
	Currency      string       `db:"currency"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}

func NewTransactionModel(transaction transaction_v1.Transaction) TransactionModel {
	return TransactionModel{
		Id:            transaction.GetId(),
		UserId:        transaction.GetUserId(),
		OperationType: transaction.GetOperationType().String(),
		Price:         transaction.GetPrice(),
		Currency:      transaction.GetCurrency().String(),
		CreatedAt:     transaction.CreatedTime.AsTime(),
	}
}

func NewTransactionModelToDto(transaction TransactionModel) transaction_v1.Transaction {
	return transaction_v1.Transaction{
		Id:            transaction.Id,
		UserId:        transaction.UserId,
		OperationType: operationTypeToEnum(transaction.OperationType),
		Price:         transaction.Price,
		Currency:      currencyToEnum(transaction.Currency),
		CreatedTime:   timestamppb.New(transaction.CreatedAt),
	}
}

func operationTypeToEnum(operationType string) transaction_v1.OperationType {
	switch strings.ToUpper(operationType) {
	case transaction_v1.OperationType_name[0]:
		return transaction_v1.OperationType_OPERATION_TYPE_CREDITING
	case transaction_v1.OperationType_name[1]:
		return transaction_v1.OperationType_OPERATION_TYPE_DEBITING

	default:
		return transaction_v1.OperationType_OPERATION_TYPE_NONE
	}
}

func currencyToEnum(currency string) transaction_v1.Currency {
	switch strings.ToUpper(currency) {
	case transaction_v1.Currency_name[0]:
		return transaction_v1.Currency_CURRENCY_USD
	case transaction_v1.Currency_name[2]:
		return transaction_v1.Currency_CURRENCY_EUR
	case transaction_v1.Currency_name[3]:
		return transaction_v1.Currency_CURRENCY_RUB

	default:
		return transaction_v1.Currency_CURRENCY_NONE
	}
}
