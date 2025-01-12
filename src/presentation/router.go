package presentation

import (
	"go-ddd-template/src/presentation/handler"
	handler_line "go-ddd-template/src/presentation/handler/line"
	"go-ddd-template/src/presentation/middleware"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, helloHandler handler.HelloHandler, accountHandler handler.AccountHandler, lineMessageHandler handler_line.LineMessageHandler, jwtMiddleware middleware.JWTMiddleware) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, go-ddd-template!")
	})

	e.GET("/hello-world/:id", helloHandler.HelloWorldDetail(), jwtMiddleware.Handle) // 認証が必要なエンドポイントにJWTミドルウェアを適用
	e.POST("/sign-in", accountHandler.SignIn())
	e.POST("/sign-up", accountHandler.SignUp())
	e.POST("/line-notify-user", lineMessageHandler.NotifyUser())

	e.GET("/slow", func(c echo.Context) error {
		// 意図的に遅延を発生させる
		time.Sleep(10 * time.Second)
		return c.String(http.StatusOK, "This won't be reached.")
	})
}
