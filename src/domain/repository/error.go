package repository

import (
	"context"
	"errors"
	"fmt"
)

type RepositoryErr error

var (
	// 無効な引数が渡された場合に使用
	ErrInvalidArgument RepositoryErr = errors.New("ErrInvalidArgument")
	// 指定されたリソースが見つからなかった場合に使用
	ErrResourceNotFound RepositoryErr = errors.New("ErrResourceNotFound")
	// リソースの競合が発生した場合に使用（例: データベースのユニーク制約違反）
	ErrResourceConflict RepositoryErr = errors.New("ErrResourceConflict")
	// 操作の前提条件が満たされていない場合に使用（例: 必須データが不整合）
	ErrFailedPrecondition RepositoryErr = errors.New("ErrFailedPrecondition")
	// リソースやサービスが現在利用できない場合に使用（例: 一時的なサーバーダウン）
	ErrUnavailable RepositoryErr = errors.New("ErrUnavailable")
	// 外部サービスが利用できない場合に使用（例: サードパーティAPIがダウン）
	ErrExternalServiceUnavailable RepositoryErr = errors.New("ErrExternalServiceUnavailable")
	// システム内部で想定外のエラーが発生した場合に使用
	ErrInternal RepositoryErr = errors.New("ErrInternal")
	// 外部サービスで内部エラーが発生した場合に使用（例: サードパーティAPIの500エラー）
	ErrExternalServiceInternal RepositoryErr = errors.New("ErrExternalServiceInternal")
	// 外部サービスで認証エラーが発生した場合に使用（例: トークンが無効または期限切れ）
	ErrExternalServiceUnauthenticated RepositoryErr = errors.New("ErrExternalServiceUnauthenticated")
	// 外部サービスでアクセス権限が拒否された場合に使用（例: 不足した権限でアクセス）
	ErrExternalServicePermissionDenied RepositoryErr = errors.New("ErrExternalServicePermissionDenied")
	
	// 操作がcontext.Canceledによってキャンセルされた場合に使用（例: タイムアウトや明示的なキャンセル）
	ErrCanceled RepositoryErr = fmt.Errorf("ErrCanceled: %w", context.Canceled)
)
