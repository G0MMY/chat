package model

type Room struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,alpha"`
}

type JoinRequest struct {
	UserId int `json:"userId" validate:"required,gt=0"`
	RoomId int `json:"roomId" validate:"required,gt=0"`
}
