package persistence

import (
	"context"
	"errors"

	"github.com/G0MMY/chat/model"
	"github.com/G0MMY/chat/util"
	"github.com/go-playground/validator/v10"
)

type UserStorable interface {
	AddUser(user *model.User) (*model.User, error)
	Login(user *model.User) (string, error)
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

func (s *userStore) Login(user *model.User) (string, error) {
	if err := s.validate.Var(user.Password, "required"); err != nil {
		return "", err
	}
	errEmail := s.validate.Var(user.Email, "required,email")
	errUsername := s.validate.Var(user.Username, "required,alpha")
	if errEmail != nil && errUsername != nil {
		return "", errors.New("the username and the email wasn't provided")
	}

	params := []interface{}{user.Password}
	query := "SELECT id, username, email FROM users WHERE password=$1"

	if errUsername == nil {
		query += " AND username=$2"
		params = append(params, user.Username)
	} else {
		query += " AND email=$2"
		params = append(params, user.Email)
	}
	var id int
	var username, email string

	err := s.connection.conn.QueryRow(context.Background(), query, params...).Scan(&id, &username, &email)
	if err != nil {
		return "", err
	}

	return util.CreateToken(&model.User{Id: id, Username: username, Email: email})
}
