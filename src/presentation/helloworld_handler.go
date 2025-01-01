package presentation

import (
	"net/http"
	"strconv"

	"go-ddd-template/src/usecase"

	"github.com/labstack/echo"
)

type HelloHandler struct {
	helloUsecase usecase.HelloWorldUsecase
}

func NewHelloHandler(helloUsecase usecase.HelloWorldUsecase) HelloHandler {
	helloWorldHandler := HelloHandler{helloUsecase: helloUsecase}
	return helloWorldHandler
}

func (handler *HelloHandler) HelloWorldDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID format"})
		}
		model, err := handler.helloUsecase.HelloWorldDetail(id)
		if err != nil {
			switch err.Error() {
			case "not results found":
				return c.JSON(http.StatusNotFound, map[string]string{"error": "not results found"})
			default:
				// その他のエラー
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			}
		}
		return c.JSON(http.StatusOK, model)
	}
}
