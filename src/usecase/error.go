package usecase

import (
	"errors"
)

type UsecaseErr error

var (
	// ErrInvalidArgument          UsecaseErr = errors.New("ErrInvalidArgument")
	ErrResourceNotFound         UsecaseErr = errors.New("ErrResourceNotFound")
	ErrResourceConflict         UsecaseErr = errors.New("ErrResourceConflict")
	// ErrFailedPrecondition       UsecaseErr = errors.New("ErrFailedPrecondition")
	// ErrUnavailable              UsecaseErr = errors.New("ErrUnavailable")
	ErrInternal                 UsecaseErr = errors.New("ErrInternal")
	// ErrPermissionDenied         UsecaseErr = errors.New("ErrPermissionDenied")
	// ErrNoAuthorization          UsecaseErr = errors.New("ErrNoAuthorization")
	// ErrInvalidAccessToken       UsecaseErr = errors.New("ErrInvalidOrExpiredAccessToken")
	// ErrInvalidRefreshToken      UsecaseErr = errors.New("ErrInvalidRefreshToken")
	// ErrInvalidAuthorizationCode UsecaseErr = errors.New("ErrInvalidAuthorizationCode")

	// // errors.Is で区別したいため context.Canceled を wrap している
	// ErrCanceled UsecaseErr = fmt.Errorf("ErrCancled: %w", context.Canceled)
)