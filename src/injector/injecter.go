package injector

import (
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"go-ddd-template/src/presentation"
	"go-ddd-template/src/usecase"
)

func InjectDB() infrastructure.SqlHandler {
	sqlhandler := infrastructure.NewSqlHandler()
	return *sqlhandler
}

func InjectHelloRepository() repository.HelloRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewHelloRepository(sqlHandler)
}

func InjectHelloUsecase() usecase.HelloWorldUsecase {
	HelloRepo := InjectHelloRepository()
	return usecase.NewHelloWorldUsecase(HelloRepo)
}

func InjectHelloHandler() presentation.HelloHandler {
	return presentation.NewHelloHandler(InjectHelloUsecase())
}
