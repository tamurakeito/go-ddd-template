package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 開発用
// bcrypt.GenerateFromPasswordでハッシュ化した文字列を生成

func main() {
	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Println("Hashed password:", string(hashedPassword))
}
