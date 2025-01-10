package repository

import (
	"context"
	"go-ddd-template/src/domain/entity"
)

type HelloRepository interface {
	Find(ctx context.Context,id int) (hello entity.Hello, err error)
}

type AccountRepository interface {
	FindUserId(ctx context.Context,userId string) (account entity.Account, err error)
	Create(ctx context.Context,userId string, password string, name string) (account entity.Account, err error)
}
