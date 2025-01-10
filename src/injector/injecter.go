package injector

import (
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"go-ddd-template/src/presentation/handler"
	"go-ddd-template/src/presentation/middleware"
	"go-ddd-template/src/service"
	"go-ddd-template/src/usecase"
	"time"
)

// InjectDB provides a SQL handler
func InjectDB() infrastructure.SqlHandler {
	sqlhandler := infrastructure.NewSqlHandler()
	return *sqlhandler
}

// Repository injection
func InjectHelloRepository() repository.HelloRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewHelloRepository(sqlHandler)
}
func InjectAccountRepository() repository.AccountRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewAccountRepository(sqlHandler)
}

// Service injection
func InjectAuthService() service.AuthService {
	secretKey := infrastructure.LoadJWTSecret()
	return service.NewAuthService(secretKey)
}
func InjectEncryptService() service.EncryptService {
	return service.NewEncryptService()
}

// Usecase injection
func InjectHelloUsecase() usecase.HelloWorldUsecase {
	helloRepo := InjectHelloRepository()
	return usecase.NewHelloWorldUsecase(helloRepo)
}
func InjectAccountUsecase() usecase.AccountUsecase {
	accountRepo := InjectAccountRepository()
	authServ := InjectAuthService()
	encryptServ := InjectEncryptService()
	return usecase.NewAccountUsecase(accountRepo, authServ, encryptServ)
}

// Handler injection
func InjectHelloHandler() handler.HelloHandler {
	return handler.NewHelloHandler(InjectHelloUsecase())
}
func InjectAccountHandler() handler.AccountHandler {
	return handler.NewAccountHandler(InjectAccountUsecase())
}

// Middleware injection
func InjectJWTMiddleware() middleware.JWTMiddleware {
	authService := InjectAuthService()
	return middleware.NewJWTMiddleware(authService)
}
func InjectCORSMiddleware() middleware.CORSMiddleware {
	return middleware.NewCORSMiddleware()
}
func InjectTimeoutMiddleware() middleware.TimeoutMiddleware {
	return middleware.NewTimeoutMiddleware(5 * time.Second)
}
