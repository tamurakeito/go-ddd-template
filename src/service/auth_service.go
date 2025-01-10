package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	GenerateToken(userId string) (string, error)
	VerifyToken(tokenString string) (map[string]interface{}, error)

}

type jwtConfig struct {
	secretKey string
}

func NewAuthService(secretKey string) AuthService {
	return &jwtConfig{secretKey: secretKey}
}

// GenerateToken はJWTトークンを生成します
func (j *jwtConfig) GenerateToken(userId string) (string, error) {
	// クレームの設定
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(72 * time.Hour).Unix(), // 72時間有効
	}

	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// VerifyToken はJWTトークンを検証します
func (j *jwtConfig) VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}