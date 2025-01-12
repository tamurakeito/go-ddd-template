package usecase

import (
	"context"
	repository_line "go-ddd-template/src/domain/repository/line"
)

type LineMessageUsecase interface {
	NotifyUser(ctx context.Context, userId string, message string) (err error) 
}

type lineMessageUsecase struct {
    lineRepo repository_line.LineMessageRepository
}

func NewLineMessageUsecase(lineRepo repository_line.LineMessageRepository) LineMessageUsecase {
	lineMessageUsecase := lineMessageUsecase{lineRepo: lineRepo}
	return &lineMessageUsecase
}

func (u *lineMessageUsecase) NotifyUser(ctx context.Context, userId string, message string) (err error) {
    err = u.lineRepo.PushMessage(ctx, userId, message)
	// エラーハンドリングを
	return
}