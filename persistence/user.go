package persistence

import (
	"context"

	"github.com/G0MMY/chat/model"
	"github.com/go-playground/validator/v10"
)

type UserStorable interface {
	AddUser(user *model.User) (*model.User, error)
}

type userStore struct {
	validate   *validator.Validate
	connection *connection
}

func NewUserStore(connection *connection) UserStorable {
	return &userStore{connection: connection, validate: validator.New()}
}

func (s *userStore) AddUser(user *model.User) (*model.User, error) {
	query := "INSERT INTO user VALUES ($1,$2,$3)"

	var id int

	err := s.connection.conn.QueryRow(context.Background(), query, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.Id = id

	return user, nil
}
