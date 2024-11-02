package usecase

import (
	"log"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
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

func (usecase *helloWorldUsecase) HelloWorldDetail(id int) (detail model.HelloWorld, err error) {
	hello, err := usecase.helloRepo.Find(id)
	if err != nil {
		log.Fatal(err)
		return detail, err
	}

	hellos := make([]model.Hello, 0)
	hellos = append(hellos, hello)

	detail = model.HelloWorld{Id: hello.Id, Hello: hellos}
	return
}
