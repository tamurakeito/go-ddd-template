package main

import (
	"fmt"

	"go-ddd-template/src/injector"
	injector_line "go-ddd-template/src/injector/line"
	"go-ddd-template/src/presentation"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	helloHandler := injector.InjectHelloHandler()
	accountHandler := injector.InjectAccountHandler()
	lineMessageHandler := injector_line.InjectLineMessageHandler()
	jwtMiddleware := injector.InjectJWTMiddleware()
	corsMiddleware := injector.InjectCORSMiddleware()
	timeoutMiddleware := injector.InjectTimeoutMiddleware()

	e := echo.New()
	// 共通のミドルウェア
	e.Use(corsMiddleware.Handle)
	e.Use(corsMiddleware.Handle,timeoutMiddleware.Handle)

	presentation.InitRouting(e, helloHandler, accountHandler, lineMessageHandler, jwtMiddleware)
	// Logger.Fatalはエラーメッセージをログに出力しアプリケーションを停止する
	// 重要なエラーが発生した場合に使用される
	// 普通のエラーは通常のエラーハンドリングを使おう
	e.Logger.Fatal(e.Start(":8080"))
}
