package listener

import (
	"github.com/sirupsen/logrus"
	"transactionService/internal/broker/event"
	"transactionService/internal/repository"
	"transactionService/internal/repository/models"
	"transactionService/pkg/proto/transaction_v1"
	"transactionService/pkg/proto/user_v1"
)

type Listener struct {
	repos *repository.Repository
}

func NewListener(repos *repository.Repository) *Listener {
	return &Listener{repos: repos}
}

func (l *Listener) Handle(event event.Event) error {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in Handle", err)
		}
	}()
	if event.UserEvent != nil {
		return l.handleUserEvent(*event.UserEvent)
	}
	if event.TransactionEvent != nil {
		return l.handleTransactionEvent(*event.TransactionEvent)
	}
	return nil
}

func (l *Listener) handleUserEvent(user user_v1.User) error {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in handleUserEvent", err)
		}
	}()
	_, err := l.repos.UsersRepository.Create(models.NewUserModel(user))
	return err
}

func (l *Listener) handleTransactionEvent(transaction transaction_v1.Transaction) error {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in handleTransactionEvent", err)
		}
	}()
	user, err := l.repos.UsersRepository.GetById(transaction.GetUserId())
	_, err = l.repos.TransactionsRepository.Create(models.NewTransactionModel(transaction))
	if err != nil {
		return err
	}
	if transaction.GetOperationType() == transaction_v1.OperationType_OPERATION_TYPE_CREDITING {
		user.Balance += transaction.GetPrice()
	}
	if transaction.GetOperationType() == transaction_v1.OperationType_OPERATION_TYPE_DEBITING {
		user.Balance -= transaction.GetPrice()
	}
	return l.repos.UsersRepository.UpdateBalance(user.Id, user.Balance)
}
