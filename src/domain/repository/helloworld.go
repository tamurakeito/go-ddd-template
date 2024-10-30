package repository

import (
	"go-ddd-template/src/domain/model"
)

type HelloRepository interface {
	Find(id int) (hello model.Hello, err error)
}
