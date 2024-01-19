package model

type Invitation struct {
	Id       int    `json:"id"`
	Sender   string `json:"sender" validate:"required,alpha"`
	Receiver string `json:"receiver" validate:"required,alpha"`
	RoomName string `json:"roomName" validate:"required,alpha"`
	RoomId   int    `json:"roomId" validate:"gt=-1"`
}
