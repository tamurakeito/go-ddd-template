package usecase_helloworld

import (
	"database/sql"
	"errors"
	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/usecase"
	"log"
)

type HelloWorldUsecase interface {
	HelloWorldDetail(id int) (detail model.HelloWorld, err error)
}

type helloWorldUsecase struct {
	helloRepo repository.HelloRepository
}

func NewHelloWorldUsecase(helloRepo repository.HelloRepository) HelloWorldUsecase {
	helloUsecase := helloWorldUsecase{helloRepo: helloRepo}
	return &helloUsecase
}

func (u *helloWorldUsecase) HelloWorldDetail(id int) (detail model.HelloWorld, err error) {
	hello, err := u.helloRepo.Find(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = usecase.ErrResultsNotFound
			return
		}
		err = usecase.ErrFailedToRetrieveData
		log.Printf("[SignIn] Error retrieving data: %v", err)
		return
	}
	detail = model.HelloWorld{Id: hello.Id, Hello: hello}
	return
}
