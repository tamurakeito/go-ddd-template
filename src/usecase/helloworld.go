package usecase

import (
	"errors"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
)

type HelloWorldUsecase interface {
	HelloWorldDetail(id int) (detail entity.HelloWorld, err error)
}

type helloWorldUsecase struct {
	helloRepo repository.HelloRepository
}

func NewHelloWorldUsecase(helloRepo repository.HelloRepository) HelloWorldUsecase {
	helloUsecase := helloWorldUsecase{helloRepo: helloRepo}
	return &helloUsecase
}

func (u *helloWorldUsecase) HelloWorldDetail(id int) (detail entity.HelloWorld, err error) {
	hello, err := u.helloRepo.Find(id)
	if err != nil {
		if errors.Is(err, repository.ErrResourceNotFound) {
			err = ErrResourceNotFound
			return
		}
		err = ErrInternal
		return
	}
	detail = entity.HelloWorld{Id: hello.Id, Hello: hello}
	return
}
