package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("ANVmOwG4K3rwvmCcetr91HPpFg")

func GerarToken(usuario string) (string, error) {
	claims := jwt.MapClaims{
		"usuario": usuario,
		"exp":     time.Now().Add(time.Minute * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidarToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return errors.New("Token inválido ou expirado")
	}

	return nil
}
