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
			Id:     account.Id,
			UserId: account.UserId,
			Name:   account.Name,
			Token:  token,
		}

		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		// }
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

		return c.JSON(http.StatusOK, result)
	}

}
