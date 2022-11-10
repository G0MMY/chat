package model

type Room struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,alpha"`
}

type JoinRequest struct {
	Username string `json:"username" validate:"required,alpha"`
	RoomId   int    `json:"roomId" validate:"required,gt=0"`
}
