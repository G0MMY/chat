package persistence

import (
	"context"
	"time"

	"github.com/G0MMY/chat/model"
	"github.com/go-playground/validator/v10"
)

type MessageStorable interface {
	AddMessage(msg *model.Message) (*model.Message, error)
	GetMessages(roomId string) (*model.Messages, error)
}

type messageStore struct {
	validate   *validator.Validate
	connection *Connection
}

func NewMessageStore(connection *Connection) MessageStorable {
	return &messageStore{connection: connection, validate: validator.New()}
}

func (s *messageStore) AddMessage(msg *model.Message) (*model.Message, error) {
	if err := s.validate.Struct(msg); err != nil {
		return nil, err
	}

	var id int
	query := "INSERT INTO messages (room_id, sender, send_time, msg) VALUES ($1,$2,$3,$4) RETURNING id"

	err := s.connection.conn.QueryRow(context.Background(), query, msg.RoomId, msg.Sender, time.Unix(msg.SendTime, 0), msg.Msg).Scan(&id)
	if err != nil {
		return nil, err
	}

	msg.Id = id

	return msg, nil
}

func (s *messageStore) GetMessages(roomId string) (*model.Messages, error) {
	if err := s.validate.Var(roomId, "required"); err != nil {
		return nil, err
	}

	query := "SELECT id, room_id, sender, send_time, msg FROM messages WHERE room_id=$1 ORDER BY send_time"

	rows, err := s.connection.conn.Query(context.Background(), query, roomId)
	if err != nil {
		return nil, err
	}

	var messages model.Messages
	for rows.Next() {
		var id, roomId int
		var sender, msg string
		var sendTime time.Time

		err := rows.Scan(&id, &roomId, &sender, &sendTime, &msg)
		if err != nil {
			return nil, err
		}

		messages.Items = append(messages.Items, model.Message{Id: id, RoomId: roomId, Sender: sender, SendTime: sendTime.Unix(), Msg: msg})
	}

	return &messages, nil
}
