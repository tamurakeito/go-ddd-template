package usecase

import (
	"errors"
)

var (
	// common

	// example
	ErrResultsNotFound = errors.New("ResultsNotFound")
	ErrFailedToRetrieveData = errors.New("FailedToRetrieveData")

	// account
	ErrUserNotFound        = errors.New("UserNotFound")
	ErrInvalidPassword     = errors.New("InvalidPassword")
	ErrFailedToGenerateToken = errors.New("FailedToGenerateToken")
	ErrFailedToRetrieveUser = errors.New("FailedToRetrieveUser")
	ErrFailedToHashPassword = errors.New("FailedToHashPassword")
	ErrFailedToCreateAccount = errors.New("FailedToCreateAccount")
)