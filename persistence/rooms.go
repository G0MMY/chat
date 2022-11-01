package persistence

import (
	"context"
	"errors"

	"github.com/G0MMY/chat/model"
	"github.com/go-playground/validator/v10"
)

type RoomStorable interface {
	GetRoom(roomName string) (*model.Room, error)
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

func (s *roomStore) GetRoom(roomName string) (*model.Room, error) {
	if err := s.validate.Var(roomName, "required,alpha"); err != nil {
		return nil, err
	}

	query := "SELECT id, name FROM rooms WHERE name=$1"

	rows, err := s.connection.conn.Query(context.Background(), query, roomName)
	if err != nil {
		return nil, err
	}

	var room *model.Room
	for rows.Next() {
		if room != nil {
			return nil, errors.New("duplicate rooms found")
		}

		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		room = &model.Room{Id: id, Name: name}
	}

	return room, nil
}

func (s *roomStore) JoinRoom(joinRequest *model.JoinRequest) error {
	if err := s.validate.Struct(joinRequest); err != nil {
		return err
	}

	query := "INSERT INTO rooms_users (room_id, user_id) VALUES ($1, $2)"

	_, err := s.connection.conn.Exec(context.Background(), query, joinRequest.RoomId, joinRequest.UserId)
	if err != nil {
		return err
	}

	return nil
}
