package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"transactionService/internal/repository/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) Create(user models.UserModel) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, $2) RETURNING id", usersTable)

	row := u.db.QueryRow(query, user.Id, user.Balance)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (u *UserPostgres) Update(user models.UserModel) error {
	// todo в будущем реализавать
	return nil
}

func (u *UserPostgres) UpdateBalance(userId string, userBalance float64) error {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in UpdateBalance", err)
		}
	}()
	query := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id=$2", usersTable)
	_, err := u.db.Exec(query, userBalance, userId)
	return err
}

func (u *UserPostgres) GetById(userId string) (models.UserModel, error) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in GetById", err)
		}
	}()
	var user models.UserModel
	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1", usersTable)
	err := u.db.Get(&user, query, userId)

	return user, err
}
