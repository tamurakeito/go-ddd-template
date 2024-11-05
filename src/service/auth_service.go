package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type JWTConfig struct {
	secretKey []byte
}

func NewAuthService(secretKey string) AuthService {
	authService := JWTConfig{secretKey: []byte(secretKey)}
	return &authService
}

func (con *JWTConfig) GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(con.secretKey)
}

// NewJWTMiddleware はJWTミドルウェアを設定して返す
func (con *JWTConfig) JWTMiddleware() echo.MiddlewareFunc {
	jwtSecret := con.secretKey
	if len(jwtSecret) == 0 {
		panic("JWT_SECRET_KEY environment variable is not set")
	}

	config := middleware.JWTConfig{
		SigningKey:  jwtSecret,
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	}

	return middleware.JWTWithConfig(config)
}
