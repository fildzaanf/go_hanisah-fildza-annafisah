package middlewares

import (
	"Praktikum/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId int, name string) (string, error) {

	payload := jwt.MapClaims{}
	payload["userId"] = userId
	payload["name"] = name
	payload["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
