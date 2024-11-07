package service

import "github.com/labstack/echo"

type AuthService interface {
	GenerateToken(userId string) (string, error)
	JWTMiddleware() echo.MiddlewareFunc
}

type EncryptService interface {
	ComparePassword(hashedPassword, password string) error
}
