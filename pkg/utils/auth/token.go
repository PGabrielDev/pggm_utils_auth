package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToke(userId string, timeValidate int) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = userId
	claims["ext"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeValidate)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("SECRET_KEY_TOKEN")
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}
