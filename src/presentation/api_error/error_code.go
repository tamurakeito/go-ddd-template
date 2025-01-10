package api_error

import "net/http"

type ApiErrorCode string

const (
	// 汎用エラー
	// client error
	// "E1xxxxxxxx" 台 クライアントエラー
	// "E100100000" 台 クライアントエラー 汎用エラー
	// "E2xxxxxxxx" 台 サーバエラー
	// "E2001xxxxx" 台 サーバエラー 汎用エラー

	InvalidArgument    = ApiErrorCode("E100100001")
	ResourceNotFound   = ApiErrorCode("E100100002")
	ResourceConflict   = ApiErrorCode("E100100003")
	PermissionDenied   = ApiErrorCode("E100100004")
	FailedPrecondition = ApiErrorCode("E100100005")
	Unauthenticated    = ApiErrorCode("E100100006")

	// server error
	Internal    = ApiErrorCode("E200100001")
	Unavailable = ApiErrorCode("E200100002") // retryable
	Timeout = ApiErrorCode("E200100003")
	// 個別エラー
)

func ApiErrorCodeToStatusCode(code ApiErrorCode) int {

	switch code {

	// client error
	case InvalidArgument:
		return http.StatusBadRequest
	case ResourceNotFound:
		return http.StatusNotFound
	case ResourceConflict:
		return http.StatusConflict
	case PermissionDenied:
		return http.StatusForbidden
	case Unauthenticated:
		return http.StatusUnauthorized
	case FailedPrecondition:
		return http.StatusPreconditionFailed

	// server error
	case Internal:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	case Timeout:
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}
