package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CORSMiddleware struct{}

func NewCORSMiddleware() CORSMiddleware {
	return CORSMiddleware{}
}

func (m CORSMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	})(next)
}
