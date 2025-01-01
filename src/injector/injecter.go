package injector

import (
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"go-ddd-template/src/presentation/handler"
	"go-ddd-template/src/service"
	usecase_account "go-ddd-template/src/usecase/account"
	usecase_helloworld "go-ddd-template/src/usecase/helloworld"
)

func InjectDB() infrastructure.SqlHandler {
	sqlhandler := infrastructure.NewSqlHandler()
	return *sqlhandler
}

func InjectHelloRepository() repository.HelloRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewHelloRepository(sqlHandler)
}
func InjectHelloUsecase() usecase_helloworld.HelloWorldUsecase {
	helloRepo := InjectHelloRepository()
	return usecase_helloworld.NewHelloWorldUsecase(helloRepo)
}
func InjectHelloHandler() handler.HelloHandler {
	return handler.NewHelloHandler(InjectHelloUsecase())
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
func InjectAccountUsecase() usecase_account.AccountUsecase {
	accountRepo := InjectAccountRepository()
	authServ := InjectAuthService()
	encryptServ := InjectEncryptService()
	return usecase_account.NewAccountUsecase(accountRepo, authServ, encryptServ)
}
func InjectAccountHandler() handler.AccountHandler {
	return handler.NewAccountHandler(InjectAccountUsecase())
}
