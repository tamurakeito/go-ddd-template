package handler

import (
	"net/http"

	"go-ddd-template/src/domain/entity"
	api_errors "go-ddd-template/src/presentation/errors"
	usecase_account "go-ddd-template/src/usecase/account"

	"github.com/labstack/echo"
)

type AccountHandler struct {
	accountUsecase usecase_account.AccountUsecase
}

func NewAccountHandler(accountUsecase usecase_account.AccountUsecase) AccountHandler {
	accountHandler := AccountHandler{accountUsecase: accountUsecase}
	return accountHandler
}

func (handler *AccountHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		body := entity.SignInRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, api_errors.ErrInvalidRequest)
		}

		account, token, err := handler.accountUsecase.SignIn(body.UserId, body.Password)
		if err != nil {
			return c.JSON(api_errors.GetErrorResponse(err))
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
			return c.JSON(http.StatusBadRequest, api_errors.ErrInvalidRequest)
		}

		account, token, err := handler.accountUsecase.SignUp(body.UserId, body.Password, body.Name)
		if err != nil {
			return c.JSON(api_errors.GetErrorResponse(err))
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
