package handler

import (
	"net/http"
	"strconv"

	api_errors "go-ddd-template/src/presentation/errors"
	usecase_helloworld "go-ddd-template/src/usecase/helloworld"

	"github.com/labstack/echo"
)

type HelloHandler struct {
	helloUsecase usecase_helloworld.HelloWorldUsecase
}

func NewHelloHandler(helloUsecase usecase_helloworld.HelloWorldUsecase) HelloHandler {
	helloWorldHandler := HelloHandler{helloUsecase: helloUsecase}
	return helloWorldHandler
}

func (handler *HelloHandler) HelloWorldDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, api_errors.ErrInvalidRequest)
		}
		model, err := handler.helloUsecase.HelloWorldDetail(id)
		if err != nil {
			return c.JSON(api_errors.GetErrorResponse(err))
		}
		return c.JSON(http.StatusOK, model)
	}
}
