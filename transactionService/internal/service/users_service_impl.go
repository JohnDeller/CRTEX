package service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"transactionService/internal/broker/event"
	"transactionService/internal/broker/publisher"
	"transactionService/internal/repository/interfaces"
	"transactionService/internal/repository/models"
	"transactionService/pkg/proto/user_v1"
)

const queueName = "event_queue"

type UsersServiceImpl struct {
	repo interfaces.UsersRepository
	conn *amqp.Connection
}

func NewUsersServiceImpl(repo interfaces.UsersRepository, conn *amqp.Connection) *UsersServiceImpl {
	return &UsersServiceImpl{repo: repo, conn: conn}
}

func (u *UsersServiceImpl) Create(user user_v1.User) (string, error) {
	user.Id = uuid.NewV4().String()
	return user.GetId(), publisher.Publish(u.conn, event.NewUserEvent(&user), queueName)
}

func (u *UsersServiceImpl) Update(user user_v1.User) error {
	return u.repo.Update(models.NewUserModel(user))
}

func (u *UsersServiceImpl) UpdateBalance(userId string, userBalance float64) error {
	return u.repo.UpdateBalance(userId, userBalance)
}

func (u *UsersServiceImpl) GetById(userId string) (user_v1.User, error) {
	user, err := u.repo.GetById(userId)
	return models.NewUserModelToDto(user), err
}
