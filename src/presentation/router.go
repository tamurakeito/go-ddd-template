package presentation

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, helloHandler HelloHandler) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, go-ddd-template!")
	})
	e.GET("/hello-world/:id", helloHandler.HelloWorldDetail())
}
