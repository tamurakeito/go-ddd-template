package injector

import (
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"go-ddd-template/src/presentation"
	"go-ddd-template/src/service"
	"go-ddd-template/src/usecase"
	"os"
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
	helloRepo := InjectHelloRepository()
	return usecase.NewHelloWorldUsecase(helloRepo)
}
func InjectHelloHandler() presentation.HelloHandler {
	return presentation.NewHelloHandler(InjectHelloUsecase())
}

func InjectAuthRepository() repository.AuthRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewAuthRepository(sqlHandler)
}
func InjectTokenGenerator() service.TokenGenerator {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return service.NewTokenGenerator(secretKey)
}
func InjectAuthUsecase() usecase.AuthUsecase {
	authRepo := InjectAuthRepository()
	tokenGen := InjectTokenGenerator()
	return usecase.NewAuthUsecase(authRepo, tokenGen)
}
func InjectAuthHandler() presentation.AuthHandler {
	return presentation.NewAuthHandler(InjectAuthUsecase())
}
