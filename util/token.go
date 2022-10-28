package util

import (
	"github.com/G0MMY/chat/model"
	"github.com/golang-jwt/jwt/v4"
)

type userClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodPS256.SigningMethodRSA, userClaims{Username: user.Username})

	// change secret
	signedToken, err := token.SignedString([]byte("test"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetUsernameFromToken(tokenString string) (string, error) {
	var claims *userClaims

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})
	if err != nil {
		return "", err
	}

	return claims.Username, nil
}
