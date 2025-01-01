package presentation_errors

import (
	"go-ddd-template/src/usecase"
	"net/http"
)

// 定義済みエラーレスポンス
var (
	ErrInvalidRequest = map[string]interface{}{
		"error":   "InvalidRequest",
		"message": "Invalid request body",
	}

	ErrResultsNotFound = map[string]interface{}{
		"error":   "ResultsNotFound",
		"message": "Results not found",
	}

	ErrFailedToRetrieveData = map[string]interface{}{
		"error":   "FailedToRetrieveData",
		"message": "Failed to retrieve data",
	}

	ErrUserNotFound = map[string]interface{}{
		"error":   "UserNotFound",
		"message": "User not found",
	}

	ErrInvalidPassword = map[string]interface{}{
		"error":   "InvalidPassword",
		"message": "Invalid password",
	}

	ErrFailedToRetrieveUser = map[string]interface{}{
		"error":   "FailedToRetrieveUser",
		"message": "Failed to retrieve user",
	}

	ErrFailedToGenerateToken = map[string]interface{}{
		"error":   "FailedToGenerateToken",
		"message": "Failed to generate token",
	}

	ErrFailedToHashPassword = map[string]interface{}{
		"error":   "FailedToHashPassword",
		"message": "Failed to hash password",
	}

	ErrFailedToCreateAccount = map[string]interface{}{
		"error":   "FailedToCreateAccount",
		"message": "Failed to create account",
	}

	ErrInternalServerError = map[string]interface{}{
		"error":   "InternalServerError",
		"message": "An unexpected error occurred",
	}
)

// HTTP ステータスコードとエラーレスポンスを紐づけた関数
func GetErrorResponse(err error) (int, map[string]interface{}) {
	switch err {
	case usecase.ErrUserNotFound:
		return http.StatusNotFound, ErrUserNotFound
	case usecase.ErrInvalidPassword:
		return http.StatusUnauthorized, ErrInvalidPassword
	case usecase.ErrResultsNotFound:
		return http.StatusNotFound, ErrResultsNotFound
	case usecase.ErrFailedToRetrieveData:
		return http.StatusInternalServerError, ErrFailedToRetrieveData
	case usecase.ErrFailedToRetrieveUser:
		return http.StatusInternalServerError, ErrFailedToRetrieveUser
	case usecase.ErrFailedToGenerateToken:
		return http.StatusInternalServerError, ErrFailedToGenerateToken
	case usecase.ErrFailedToHashPassword:
		return http.StatusInternalServerError, ErrFailedToHashPassword
	case usecase.ErrFailedToCreateAccount:
		return http.StatusInternalServerError, ErrFailedToCreateAccount
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}
