package middleware

import (
	"net/http"

	"go-ddd-template/src/service"

	"github.com/labstack/echo"
)

type JWTMiddleware struct {
	authService service.AuthService
}

func NewJWTMiddleware(authService service.AuthService) JWTMiddleware {
	return JWTMiddleware{authService: authService}
}

func (m JWTMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Authorization ヘッダーからトークンを取得
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || len(authHeader) <= len("Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing or invalid token"})
		}

		// "Bearer " を取り除く
		tokenString := authHeader[len("Bearer "):]

		// トークンを検証
		claims, err := m.authService.VerifyToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		// トークンのクレームをコンテキストにセット
		c.Set("userId", claims["userId"])
		return next(c)
	}
}
