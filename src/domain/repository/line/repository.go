package repository

import "context"

type LineMessageRepository interface {
    PushMessage(ctx context.Context, userId string, message string) error
}