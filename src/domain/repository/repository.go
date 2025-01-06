package repository

import (
	"go-ddd-template/src/domain/entity"
)

type HelloRepository interface {
	Find(id int) (hello entity.Hello, err error)
}

type AccountRepository interface {
	FindUserId(userId string) (account entity.Account, err error)
	Create(userId string, password string, name string) (account entity.Account, err error)
}
