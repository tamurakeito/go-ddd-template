package injector

import (
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"go-ddd-template/src/presentation"
	"go-ddd-template/src/service"
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
func InjectAuthService() service.AuthService {
	secretKey := infrastructure.LoadJWTSecret()
	return service.NewAuthService(secretKey)
}
func InjectEncryptService() service.EncryptService {
	return service.NewEncryptService()
}
func InjectAuthUsecase() usecase.AuthUsecase {
	authRepo := InjectAuthRepository()
	authServ := InjectAuthService()
	encryptServ := InjectEncryptService()
	return usecase.NewAuthUsecase(authRepo, authServ, encryptServ)
}
func InjectAuthHandler() presentation.AuthHandler {
	return presentation.NewAuthHandler(InjectAuthUsecase())
}
