package infrastructure

import (
	"database/sql"
	"log"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
)

type AuthRepository struct {
	SqlHandler
}

func NewAuthRepository(sqlHandler SqlHandler) repository.AuthRepository {
	authRepository := AuthRepository{sqlHandler}
	return &authRepository
}

func (authRepo *AuthRepository) FindUserId(userId string) (account model.Account, err error) {
	row := authRepo.SqlHandler.Conn.QueryRow("SELECT id, user_id, password, name FROM accounts WHERE user_id = ?", userId)
	err = row.Scan(&account.Id, &account.UserId, &account.Password, &account.Name)

	if err == sql.ErrNoRows {
		// データが存在しない場合の処理
		log.Println("No account found for userId:", userId)
		return account, nil // nilを返すことでエラーとしない
	} else if err != nil {
		log.Fatal(err)
		return
	}
	return
}
