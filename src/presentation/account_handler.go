package presentation

import (
	"net/http"

	"go-ddd-template/src/domain/model"
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
		body := model.SignInRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "invalid_request",
				"message": "Invalid request body",
			})
		}

		account, token, err := handler.accountUsecase.SignIn(body.UserId, body.Password)

		if err != nil {
			switch err.Error() {
			case "user not found":
				// ユーザーが存在しない場合のレスポンス
				return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			case "invalid password":
				// パスワードが一致しない場合のレスポンス
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid password"})
			default:
				// その他のエラー
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			}
			// switch {
			// case errors.Is(err, usecase.ErrUserNotFound):
			// 	return c.JSON(http.StatusNotFound, map[string]interface{}{
			// 		"error":   "user_not_found",
			// 		"message": "User not found",
			// 	})
			// case errors.Is(err, usecase.ErrInvalidPassword):
			// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			// 		"error":   "invalid_password",
			// 		"message": "Invalid password",
			// 	})
			// default:
			// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			// 		"error":   "internal_server_error",
			// 		"message": "An unexpected error occurred",
			// 	})
			// }
		}

		result := model.SignInResponse{
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
		body := model.SignUpRequest{}

		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		account, token, err := handler.accountUsecase.SignUp(body.UserId, body.Password, body.Name)

		if err != nil {
			switch err.Error() {
			case "user not found":
				// ユーザーが存在しない場合のレスポンス
				return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			case "invalid password":
				// パスワードが一致しない場合のレスポンス
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid password"})
			default:
				// その他のエラー
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			}
		}

		result := model.SignUpResponse{
			Id:     account.Id,
			UserId: account.UserId,
			Name:   account.Name,
			Token:  token,
		}

		return c.JSON(http.StatusOK, result)
	}
}
