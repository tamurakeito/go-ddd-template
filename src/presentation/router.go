package presentation

import (
	"go-ddd-template/src/service"
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, helloHandler HelloHandler, authHandler AuthHandler, authService service.AuthService) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, go-ddd-template!")
	})

	e.GET("/hello-world/:id", helloHandler.HelloWorldDetail(), authService.JWTMiddleware()) // 認証が必要なエンドポイントにJWTミドルウェアを適用
	e.POST("/sign-in", authHandler.SignIn())
}
