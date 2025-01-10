package handler

import (
	"errors"
	"net/http"

	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/presentation/api_error"
	"go-ddd-template/src/usecase"

	"github.com/labstack/echo"
)

type AccountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(accountUsecase usecase.AccountUsecase) AccountHandler {
	accountHandler := AccountHandler{accountUsecase: accountUsecase}
	return accountHandler
}

func (handler *AccountHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		body := entity.SignInRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(api_error.NewInvalidArgumentError(err))
		}

		account, token, err := handler.accountUsecase.SignIn(ctx, body.UserId, body.Password)
		if err != nil {
			if errors.Is(err, usecase.ErrDatabaseUnavailable) {
				return c.JSON(api_error.NewUnavailableError(err))
			} else if errors.Is(err, usecase.ErrResourceNotFound) {
				return c.JSON(api_error.NewResourceNotFoundError(err))
			} else {
				return c.JSON(api_error.NewInternalError(err))
			}
		}

		result := entity.SignInResponse{
			Id:     account.Id,
			UserId: account.UserId,
			Name:   account.Name,
			Token:  token,
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (handler *AccountHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		body := entity.SignUpRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(api_error.NewInvalidArgumentError(err))
		}

		account, token, err := handler.accountUsecase.SignUp(ctx, body.UserId, body.Password, body.Name)
		if err != nil {
			if errors.Is(err, usecase.ErrDatabaseUnavailable) {
				return c.JSON(api_error.NewUnavailableError(err))
			} else if errors.Is(err, usecase.ErrResourceConflict) {
				return c.JSON(api_error.NewResourceConflictError(err))
			} else {
				return c.JSON(api_error.NewInternalError(err))
			}
		}

		result := entity.SignUpResponse{
			Id:     account.Id,
			UserId: account.UserId,
			Name:   account.Name,
			Token:  token,
		}
		return c.JSON(http.StatusOK, result)
	}
}
