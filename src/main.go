package main

import (
	"fmt"

	"go-ddd-template/src/injector"
	handler "go-ddd-template/src/presentation"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("sever start")
	memoHandler := injector.InjectHelloHandler()
	authHandler := injector.InjectAuthHandler()
	authService := injector.InjectAuthService()
	e := echo.New()
	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	handler.InitRouting(e, memoHandler, authHandler, authService)
	// Logger.Fatalはエラーメッセージをログに出力しアプリケーションを停止する
	// 重要なエラーが発生した場合に使用される
	// 普通のエラーは通常のエラーハンドリングを使おう
	e.Logger.Fatal(e.Start(":8080"))
}
