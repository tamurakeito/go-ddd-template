package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

// sqlHandlerのインスタンスを生成
func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(go-ddd-template-app-db:3306)/go-ddd-template_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil
	}

	// DB接続確認
	if err := conn.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
