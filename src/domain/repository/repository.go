package repository

import (
	"go-ddd-template/src/domain/model"
)

type HelloRepository interface {
	Find(id int) (hello model.Hello, err error)
}

type AccountRepository interface {
	FindUserId(userId string) (account model.Account, err error)
	Create(userId string, password string, name string) (account model.Account, err error)
}
