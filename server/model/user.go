package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Usernames struct {
	Items []string `json:"items"`
}
