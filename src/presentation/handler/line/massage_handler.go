package handler

import (
	entity_line "go-ddd-template/src/domain/entity/line"
	"go-ddd-template/src/presentation/api_error"
	usecase_line "go-ddd-template/src/usecase/line"
	"net/http"

	"github.com/labstack/echo"
)

type LineMessageHandler struct {
	lineMessageUsecase usecase_line.LineMessageUsecase
}

func NewLinMessageHandler(lineMassageUsecase usecase_line.LineMessageUsecase) LineMessageHandler {
	lineMessageHandler := LineMessageHandler{lineMessageUsecase: lineMassageUsecase}
	return lineMessageHandler
}

func (handler * LineMessageHandler) NotifyUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		body := entity_line.NotifyUserRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(api_error.NewInvalidArgumentError(err))
		}

		err := handler.lineMessageUsecase.NotifyUser(ctx, body.UserId, body.Message)
		if err != nil {
			// エラーハンドリング
			return c.JSON(api_error.NewInternalError(err))
		}
		return c.JSON(http.StatusOK, nil)
	}
}