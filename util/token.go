package util

import (
	"io/ioutil"

	"github.com/G0MMY/chat/model"
	"github.com/golang-jwt/jwt/v4"
)

type userClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodPS256.SigningMethodRSA, userClaims{Username: user.Username})

	privateKey, err := ioutil.ReadFile("./rsa")
	if err != nil {
		return "", err
	}

	// change secret
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		privateKey, err := ioutil.ReadFile("./rsa")
		if err != nil {
			return "", err
		}

		// change secret
		key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			return "", err
		}

		return key, nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

// func GetUsernameFromToken(tokenString string) (string, error) {
// 	var claims *userClaims

// 	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("test"), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return claims.Username, nil
// }
