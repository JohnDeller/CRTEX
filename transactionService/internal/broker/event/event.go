package event

import (
	"transactionService/pkg/proto/transaction_v1"
	"transactionService/pkg/proto/user_v1"
)

type Event struct {
	UserEvent        *user_v1.User               `json:"user_event,omitempty"`
	TransactionEvent *transaction_v1.Transaction `json:"transaction_event,omitempty"`
}

func NewUserEvent(user *user_v1.User) Event {
	return Event{
		UserEvent: user,
	}
}

func NewTransactionEvent(transaction *transaction_v1.Transaction) Event {
	return Event{
		TransactionEvent: transaction,
	}
}
