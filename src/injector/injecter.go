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

func InjectAccountRepository() repository.AccountRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewAccountRepository(sqlHandler)
}
func InjectAuthService() service.AuthService {
	secretKey := infrastructure.LoadJWTSecret()
	return service.NewAuthService(secretKey)
}
func InjectEncryptService() service.EncryptService {
	return service.NewEncryptService()
}
func InjectAccountUsecase() usecase.AccountUsecase {
	accountRepo := InjectAccountRepository()
	authServ := InjectAuthService()
	encryptServ := InjectEncryptService()
	return usecase.NewAccountUsecase(accountRepo, authServ, encryptServ)
}
func InjectAccountHandler() presentation.AccountHandler {
	return presentation.NewAccountHandler(InjectAccountUsecase())
}
