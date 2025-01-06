package api_error

import "errors"

type ApiError struct {
	Err     error
	Message string
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}

// 無効な引数
func NewInvalidArgumentError(msg string) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(InvalidArgument),&ApiError{
		Err:     errors.New(msg),
		Message: msg,
	}
}

// リソースが見つからない
func NewResourceNotFoundError(msg string) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(ResourceNotFound),&ApiError{
		Err:     errors.New(msg),
		Message: msg,
	}
}

// リソースの競合
func NewResourceConflictError(msg string) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(ResourceConflict),&ApiError{
		Err:     errors.New(msg),
		Message: msg,
	}
}

// // 権限拒否
// func NewFailedPermissionDeniedError(msg string) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(PermissionDenied),&ApiError{
// 		Err:     errors.New(msg),
// 		Message: msg,
// 	}
// }

// // 前提条件未達
// func NewFailedPreconditionError(msg string) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(FailedPrecondition),&ApiError{
// 		Err:     errors.New(msg),
// 		Message: msg,
// 	}
// }

// // 認証エラー
// func NewUnauthenticatedError(msg string) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(Unauthenticated),&ApiError{
// 		Err:     errors.New(msg),
// 		Message: msg,
// 	}
// }

func NewInternalError(err error, msg string) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(Internal),&ApiError{
		Err:     err,
		Message: msg,
	}
}

// // サーバーや外部サービスが一時的にダウンしている、または負荷が高すぎて応答できない場合に使用
// func NewUnavailableError(err error, msg string) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(Unavailable),&ApiError{
// 		Err:     err,
// 		Message: msg,
// 	}
// }
