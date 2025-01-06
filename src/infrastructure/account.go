package infrastructure

import (
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
)

type AccountRepository struct {
	SqlHandler
}

func NewAccountRepository(sqlHandler SqlHandler) repository.AccountRepository {
	accountRepository := AccountRepository{sqlHandler}
	return &accountRepository
}

func (accountRepo *AccountRepository) FindUserId(userId string) (account entity.Account, err error) {
	row := accountRepo.SqlHandler.Conn.QueryRow("SELECT id, user_id, password, name FROM accounts WHERE user_id = ?", userId)
	err = row.Scan(&account.Id, &account.UserId, &account.Password, &account.Name)
	return
}

func (accountRepo *AccountRepository) Create(userId string, password string, name string) (account entity.Account, err error) {
	result, err := accountRepo.SqlHandler.Conn.Exec("INSERT accounts(user_id, password, name) VALUES (?, ?, ?)", userId, password, name)
	if err != nil {
		return account, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return account, err
	}
	account = entity.Account{
		Id:       int(lastInsertId),
		UserId:   userId,
		Password: password,
		Name:     name,
	}
	return account, nil
}
