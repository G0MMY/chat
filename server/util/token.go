package util

import (
	"time"

	"github.com/G0MMY/chat/model"
	"github.com/dgrijalva/jwt-go"
)

type userClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(user *model.User) (string, error) {
	claims := &userClaims{
		Username: user.Username,
		Id:       user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("test"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (bool, error) {
	claims := &userClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})

	if err != nil {
		return false, err
	} else if claims.ExpiresAt < time.Now().Unix() {
		return false, nil
	}

	return token.Valid, nil
}

func GetInfoFromToken(tokenString string) (string, int, error) {
	claims := &userClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})
	if err != nil {
		return "", 0, err
	}

	return claims.Username, claims.Id, nil
}
