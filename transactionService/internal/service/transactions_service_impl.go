package service

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"transactionService/internal/broker/event"
	"transactionService/internal/broker/publisher"
	"transactionService/internal/repository/interfaces"
	"transactionService/internal/repository/models"
	"transactionService/pkg/proto/transaction_v1"
)

type TransactionsServiceImpl struct {
	repo     interfaces.TransactionsRepository
	userRepo interfaces.UsersRepository
	conn     *amqp.Connection
}

func NewTransactionsServiceImpl(
	repo interfaces.TransactionsRepository,
	userRepo interfaces.UsersRepository,
	conn *amqp.Connection,
) *TransactionsServiceImpl {
	return &TransactionsServiceImpl{repo: repo, userRepo: userRepo, conn: conn}
}

func (t *TransactionsServiceImpl) Create(transaction transaction_v1.Transaction) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in Create", err)
		}
	}()
	user, err := t.userRepo.GetById(transaction.GetUserId())
	if err != nil {
		return "", err
	}
	if user.Id != transaction.GetUserId() {
		return "", fmt.Errorf("user not exist")
	}
	if transaction.GetOperationType() == transaction_v1.OperationType_OPERATION_TYPE_DEBITING && user.Balance < transaction.GetPrice() {
		return "", fmt.Errorf("error, the user's balance is below the transaction amount")
	}

	transaction.Id = uuid.NewV4().String()
	return transaction.GetId(), publisher.Publish(t.conn, event.NewTransactionEvent(&transaction), queueName)
}

func (t *TransactionsServiceImpl) GetById(transactionId string) (transaction_v1.Transaction, error) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in GetById", err)
		}
	}()
	trans, err := t.repo.GetById(transactionId)
	return models.NewTransactionModelToDto(trans), err
}
