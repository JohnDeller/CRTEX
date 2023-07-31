package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"transactionService/internal/repository/models"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (t *TransactionPostgres) Create(transaction models.TransactionModel) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in Create", err)
		}
	}()
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, operation_type, price, currency, created_at) values ($1, $2, $3, $4, $5, $6) RETURNING id", transactionsTable)

	row := t.db.QueryRow(query, transaction.Id, transaction.UserId, transaction.OperationType, transaction.Price, transaction.Currency, transaction.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (t *TransactionPostgres) GetById(transactionId string) (models.TransactionModel, error) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in GetById", err)
		}
	}()
	var transaction models.TransactionModel
	query := fmt.Sprintf("SELECT id, user_id, operation_type, price, currency, created_at FROM %s WHERE id=$1", transactionsTable)
	err := t.db.Get(&transaction, query, transactionId)

	return transaction, err
}
