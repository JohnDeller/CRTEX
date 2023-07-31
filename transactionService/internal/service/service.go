package service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"transactionService/internal/repository"
	"transactionService/internal/service/interfaces"
)

type Service struct {
	interfaces.TransactionsService
	interfaces.UsersService
}

func NewService(repos *repository.Repository, conn *amqp.Connection) *Service {
	return &Service{
		TransactionsService: NewTransactionsServiceImpl(repos.TransactionsRepository, repos.UsersRepository, conn),
		UsersService:        NewUsersServiceImpl(repos.UsersRepository, conn),
	}
}
