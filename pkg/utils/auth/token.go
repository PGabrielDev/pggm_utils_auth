package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	token_secret_key       = os.Getenv("SECRET_KEY_TOKEN")
	token_secret_key_bytes = []byte(token_secret_key)
)

func GenerateToke(userId string, timeValidate int) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = userId
	claims["ext"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeValidate)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString(token_secret_key_bytes)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetToken(w http.ResponseWriter, r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		http.Error(w, "Token de autenticação não encontrado", http.StatusUnauthorized)
		return ""
	}

	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Token de autenticação inválido", http.StatusUnauthorized)
		return ""
	}

	token := parts[1]

	return token
}

func VerifyToken(token string) (*jwt.MapClaims, error) {
	var claims jwt.MapClaims
	t, err := jwt.ParseWithClaims(token, &claims, GetKey)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Erro ao verificar assinatura do token")
		} else {
			return nil, errors.New("ao analizar token")
		}
	}
	if !t.Valid {
		return nil, errors.New("token expirado")
	}
	return &claims, nil
}

func GetKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
	}
	return token, nil
}
