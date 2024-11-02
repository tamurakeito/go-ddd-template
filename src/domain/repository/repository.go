package repository

import (
	"go-ddd-template/src/domain/model"
)

type HelloRepository interface {
	Find(id int) (hello model.Hello, err error)
}

type AuthRepository interface {
	FindUserId(userId string) (account model.Account, err error)
}
