package presentation

import (
	"net/http"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/usecase"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	helloWorldHandler := AuthHandler{authUsecase: authUsecase}
	return helloWorldHandler
}

func (handler *AuthHandler) SignIn() echo.HandlerFunc {

	return func(c echo.Context) error {
		body := model.AuthRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		account, token, err := handler.authUsecase.SignIn(body.UserId, body.Password)

		result := model.AuthResponse{
			Account: account,
			Token:   token,
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}

}
