package main

import (
	"fmt"

	"go-ddd-template/src/injector"
	"go-ddd-template/src/presentation"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("sever start")
	helloHandler := injector.InjectHelloHandler()
	accountHandler := injector.InjectAccountHandler()
	authService := injector.InjectAuthService()
	e := echo.New()
	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	presentation.InitRouting(e, helloHandler, accountHandler, authService)
	// Logger.Fatalはエラーメッセージをログに出力しアプリケーションを停止する
	// 重要なエラーが発生した場合に使用される
	// 普通のエラーは通常のエラーハンドリングを使おう
	e.Logger.Fatal(e.Start(":8080"))
}
