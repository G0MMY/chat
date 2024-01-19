package persistence

import (
	"context"

	"github.com/G0MMY/chat/model"
	"github.com/go-playground/validator/v10"
)

type InvitationStorable interface {
	AddInvitation(invitation *model.Invitation) (*model.Invitation, error)
	GetInvitationsOfUser(username string) ([]model.Invitation, error)
	DeleteInvitation(id string) error
}

type invitationStore struct {
	validate   *validator.Validate
	connection *Connection
}

func NewInvitationStore(connection *Connection) InvitationStorable {
	return &invitationStore{connection: connection, validate: validator.New()}
}

func (s *invitationStore) DeleteInvitation(id string) error {
	query := "DELETE FROM invitations where id=$1"

	_, err := s.connection.conn.Query(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *invitationStore) AddInvitation(invitation *model.Invitation) (*model.Invitation, error) {
	if err := s.validate.Struct(invitation); err != nil {
		return nil, err
	}

	var id, roomId int
	query := "INSERT INTO invitations (sender, receiver, room_id) VALUES ($1,$2,(SELECT id from rooms JOIN rooms_users ON id=room_id WHERE username=$3 AND name=$4)) RETURNING id, room_id"

	err := s.connection.conn.QueryRow(context.Background(), query, invitation.Sender, invitation.Receiver, invitation.Sender, invitation.RoomName).Scan(&id, &roomId)
	if err != nil {
		return nil, err
	}

	invitation.Id = id
	invitation.RoomId = roomId

	return invitation, nil
}

func (s *invitationStore) GetInvitationsOfUser(username string) ([]model.Invitation, error) {
	if err := s.validate.Var(username, "required,alpha"); err != nil {
		return nil, err
	}

	query := "SELECT i.id,i.sender,i.receiver,i.room_id,r.name FROM invitations as i JOIN rooms as r ON i.room_id=r.id WHERE receiver=$1"

	rows, err := s.connection.conn.Query(context.Background(), query, username)
	if err != nil {
		return nil, err
	}

	var invitations []model.Invitation
	for rows.Next() {
		var id, roomId int
		var sender, receiver, roomName string

		err := rows.Scan(&id, &sender, &receiver, &roomId, &roomName)
		if err != nil {
			return nil, err
		}

		invitations = append(invitations, model.Invitation{Id: id, Sender: sender, Receiver: receiver, RoomId: roomId, RoomName: roomName})
	}

	return invitations, nil
}
