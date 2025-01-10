package usecase

import (
	"context"
	"errors"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
)

type HelloWorldUsecase interface {
	HelloWorldDetail(ctx context.Context,id int) (detail entity.HelloWorld, err error)
}

type helloWorldUsecase struct {
	helloRepo repository.HelloRepository
}

func NewHelloWorldUsecase(helloRepo repository.HelloRepository) HelloWorldUsecase {
	helloUsecase := helloWorldUsecase{helloRepo: helloRepo}
	return &helloUsecase
}

func (u *helloWorldUsecase) HelloWorldDetail(ctx context.Context,id int) (detail entity.HelloWorld, err error) {
	hello, err := u.helloRepo.Find(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrDatabaseUnavailable) {
			err = ErrDatabaseUnavailable
			return
		} else if errors.Is(err, repository.ErrResourceNotFound) {
			err = ErrResourceNotFound
			return
		}
		err = ErrInternal
		return
	}
	detail = entity.HelloWorld{Id: hello.Id, Hello: hello}
	return
}
