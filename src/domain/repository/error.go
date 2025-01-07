package repository

import (
	"errors"
)

type RepositoryErr error

var (
	// // 無効な引数
	// ErrInvalidArgument RepositoryErr = errors.New("ErrInvalidArgument")

	// 指定されたリソースがない
	ErrResourceNotFound RepositoryErr = errors.New("ErrResourceNotFound")

	// リソースの競合が発生した場合（データベースのユニーク制約違反）
	ErrResourceConflict RepositoryErr = errors.New("ErrResourceConflict")

	// // 操作の前提条件が満たされていない場合（必須データが不整合）
	// ErrFailedPrecondition RepositoryErr = errors.New("ErrFailedPrecondition")

	// // リソースやサービスが現在利用できない（一時的なサーバーダウン）
	// ErrUnavailable RepositoryErr = errors.New("ErrUnavailable")

	// // 外部サービスが利用できない（サードパーティAPIがダウン）
	// ErrExternalServiceUnavailable RepositoryErr = errors.New("ErrExternalServiceUnavailable")

	// データベースが利用できない
	ErrDatabaseUnavailable RepositoryErr = errors.New("ErrDatabaseUnavailable")

	// システム内部で想定外のエラー
	ErrInternal RepositoryErr = errors.New("ErrInternal")

	// // 外部サービスで内部エラー（サードパーティAPIの500エラー）
	// ErrExternalServiceInternal RepositoryErr = errors.New("ErrExternalServiceInternal")

	// // 外部サービスで認証エラー（トークンが無効または期限切れ）
	// ErrExternalServiceUnauthenticated RepositoryErr = errors.New("ErrExternalServiceUnauthenticated")

	// // 外部サービスでアクセス権限が拒否（不足した権限でアクセス）
	// ErrExternalServicePermissionDenied RepositoryErr = errors.New("ErrExternalServicePermissionDenied")

	// // 操作がcontext.Canceledによってキャンセルされた場合に使用（タイムアウトや明示的なキャンセル）
	// ErrCanceled RepositoryErr = fmt.Errorf("ErrCanceled: %w", context.Canceled)
)
