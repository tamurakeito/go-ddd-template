package middleware

import (
	"context"
	"go-ddd-template/src/presentation/api_error"
	"time"

	"github.com/labstack/echo"
)

type TimeoutMiddleware struct {
	Timeout time.Duration
}

func NewTimeoutMiddleware(timeout time.Duration) TimeoutMiddleware {
	return TimeoutMiddleware{Timeout: timeout}
}

func (m TimeoutMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// タイムアウト付きContextを生成
		ctx, cancel := context.WithTimeout(c.Request().Context(), m.Timeout)
		defer cancel()

		// Contextをリクエストに設定
		c.SetRequest(c.Request().WithContext(ctx))

		// ハンドラーの実行
		errCh := make(chan error, 1)
		go func() {
			errCh <- next(c)
		}()

		select {
		case <-ctx.Done(): // タイムアウト発生
			// ctx.Err() が DeadlineExceeded の場合にのみタイムアウトレスポンスを返す
			if ctx.Err() == context.DeadlineExceeded {
				return c.JSON(api_error.NewTimeoutError(ctx.Err()))
			}
			return c.JSON(api_error.NewInternalError(ctx.Err()))
		case err := <-errCh: // ハンドラーからのエラー
			return err
		}
	}
}

