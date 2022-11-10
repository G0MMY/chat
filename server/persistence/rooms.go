package persistence

import (
	"context"

	"github.com/G0MMY/chat/model"
	"github.com/go-playground/validator/v10"
)

type RoomStorable interface {
	GetRoomUsers(id string) (*model.Usernames, error)
	AddRoom(room *model.Room) (*model.Room, error)
	JoinRoom(joinRequest *model.JoinRequest) error
}

type roomStore struct {
	validate   *validator.Validate
	connection *Connection
}

func NewRoomStore(connection *Connection) RoomStorable {
	return &roomStore{connection: connection, validate: validator.New()}
}

func (s *roomStore) AddRoom(room *model.Room) (*model.Room, error) {
	if err := s.validate.Struct(room); err != nil {
		return nil, err
	}

	var id int
	query := "INSERT INTO rooms (name) VALUES ($1) RETURNING id"

	err := s.connection.conn.QueryRow(context.Background(), query, room.Name).Scan(&id)
	if err != nil {
		return nil, err
	}

	room.Id = id

	return room, nil
}

func (s *roomStore) GetRoomUsers(id string) (*model.Usernames, error) {
	if err := s.validate.Var(id, "required"); err != nil {
		return nil, err
	}

	query := "SELECT username FROM rooms_users WHERE room_id=$1"

	rows, err := s.connection.conn.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}

	var usernames model.Usernames
	for rows.Next() {
		var name string

		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}

		usernames.Items = append(usernames.Items, name)
	}

	return &usernames, nil
}

func (s *roomStore) JoinRoom(joinRequest *model.JoinRequest) error {
	if err := s.validate.Struct(joinRequest); err != nil {
		return err
	}

	query := "INSERT INTO rooms_users (room_id, username) VALUES ($1, $2)"

	_, err := s.connection.conn.Exec(context.Background(), query, joinRequest.RoomId, joinRequest.Username)
	if err != nil {
		return err
	}

	return nil
}
