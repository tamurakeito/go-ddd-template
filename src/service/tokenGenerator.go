package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGenerator struct {
	secretKey []byte
}

func NewTokenGenerator(secretKey string) TokenGenerator {
	tokenGenerator := JWTTokenGenerator{secretKey: []byte(secretKey)}
	return &tokenGenerator
}

func (gen *JWTTokenGenerator) GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(gen.secretKey)
}
