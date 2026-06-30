package helpers

import (
	"time"

	"github.com/HRitsFadhila/golang-api-wallet/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(email string) string{
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject: email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	return token
}