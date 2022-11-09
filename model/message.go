package model

type Message struct {
	Id       int    `json:"id"`
	RoomId   int    `json:"roomId" validate:"required,gt=0"`
	Sender   string `json:"sender" validate:"required,alpha"`
	SendTime int64  `json:"sendTime" validate:"required,gt=0"`
	Msg      string `json:"msg" validate:"required"`
}

type Messages struct {
	Items []Message `json:"items"`
}
