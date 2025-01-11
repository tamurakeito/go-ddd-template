package config

import (
	"log"
	"os"
)

// LoadJWTSecret は環境変数からJWT_SECRET_KEYを読み込みます
func LoadJWTSecret() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
	return secretKey
}
