package service

import "github.com/labstack/echo"

type AuthService interface {
	GenerateToken(userId string) (string, error)
	JWTMiddleware() echo.MiddlewareFunc
}
