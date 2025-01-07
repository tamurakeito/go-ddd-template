package handler

import (
	"errors"
	"net/http"
	"strconv"

	"go-ddd-template/src/presentation/api_error"
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
			return c.JSON(api_error.NewInvalidArgumentError(err))
		}
		entity, err := handler.helloUsecase.HelloWorldDetail(id)
		if err != nil {
			if errors.Is(err, usecase.ErrDatabaseUnavailable) {
				return c.JSON(api_error.NewUnavailableError(err))
			} else if errors.Is(err, usecase.ErrResourceNotFound) {
				return c.JSON(api_error.NewResourceNotFoundError(err))
			} else {
				return c.JSON(api_error.NewInternalError(err))
			}
		}
		return c.JSON(http.StatusOK, entity)
	}
}
