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
		body := entity.SignInRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(api_error.NewInvalidArgumentError(err.Error()))
		}

		account, token, err := handler.accountUsecase.SignIn(body.UserId, body.Password)
		if err != nil {
			if errors.Is(err, usecase.ErrResourceNotFound) {
				return c.JSON(api_error.NewResourceNotFoundError(err.Error()))
			} else {
				return c.JSON(api_error.NewInternalError(err, err.Error()))
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
		body := entity.SignUpRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(api_error.NewInvalidArgumentError(err.Error()))
		}

		account, token, err := handler.accountUsecase.SignUp(body.UserId, body.Password, body.Name)
		if err != nil {
			if errors.Is(err, usecase.ErrResourceConflict) {
				return c.JSON(api_error.NewResourceConflictError(err.Error()))
			} else {
				return c.JSON(api_error.NewInternalError(err, err.Error()))
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
