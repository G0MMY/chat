package model

type Invitation struct {
	Id       int    `json:"id"`
	Sender   string `json:"sender" validate:"required,alpha"`
	Receiver string `json:"receiver" validate:"required,alpha"`
	RoomId   int    `json:"roomId" validate:"required,gt=0"`
}
