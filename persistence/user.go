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
	connection *Connection
}

func NewUserStore(connection *Connection) UserStorable {
	return &userStore{connection: connection, validate: validator.New()}
}

func (s *userStore) AddUser(user *model.User) (*model.User, error) {
	if err := s.validate.Struct(user); err != nil {
		return nil, err
	}

	var id int
	query := "INSERT INTO users (username, email, password) VALUES ($1,$2,$3) RETURNING id"

	err := s.connection.conn.QueryRow(context.Background(), query, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.Id = id

	return user, nil
}
