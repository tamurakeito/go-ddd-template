package api_error

type ApiError struct {
	Err     error
	Message string
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}

// 無効な引数
func NewInvalidArgumentError(err error) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(InvalidArgument), &ApiError{
		Err:     err,
		Message: err.Error(),
	}
}

// リソースが見つからない
func NewResourceNotFoundError(err error) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(ResourceNotFound), &ApiError{
		Err:     err,
		Message: err.Error(),
	}
}

// リソースの競合
func NewResourceConflictError(err error) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(ResourceConflict), &ApiError{
		Err:     err,
		Message: err.Error(),
	}
}

// // 権限拒否
// func NewFailedPermissionDeniedError(err error) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(PermissionDenied),&ApiError{
// 		Err:     err,
// 		Message: err.Error(),
// 	}
// }

// // 前提条件未達
// func NewFailedPreconditionError(err error) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(FailedPrecondition),&ApiError{
// 		Err:     err,
// 		Message: err.Error(),
// 	}
// }

// // 認証エラー
// func NewUnauthenticatedError(err error) (int, *ApiError) {
// 	return ApiErrorCodeToStatusCode(Unauthenticated),&ApiError{
// 		Err:     err,
// 		Message: err.Error(),
// 	}
// }

func NewInternalError(err error) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(Internal), &ApiError{
		Err:     err,
		Message: err.Error(),
	}
}

// サーバーや外部サービスが一時的にダウンしている、または負荷が高すぎて応答できない場合に使用
func NewUnavailableError(err error) (int, *ApiError) {
	return ApiErrorCodeToStatusCode(Unavailable), &ApiError{
		Err:     err,
		Message: err.Error(),
	}
}
