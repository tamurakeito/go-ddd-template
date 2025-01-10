package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"log"
)

type AccountRepository struct {
	SqlHandler
}

func NewAccountRepository(sqlHandler SqlHandler) repository.AccountRepository {
	accountRepository := AccountRepository{sqlHandler}
	return &accountRepository
}

func (accountRepo *AccountRepository) checkConnection() error {
	if accountRepo.SqlHandler.Conn == nil {
		log.Printf("[Error]AccountRepository: Database connection is nil")
		return repository.ErrDatabaseUnavailable
	}
	return nil
}

func (accountRepo *AccountRepository) FindUserId(ctx context.Context,userId string) (account entity.Account, err error) {
	if err = accountRepo.checkConnection(); err != nil {
		return
	}
	row := accountRepo.SqlHandler.Conn.QueryRow("SELECT id, user_id, password, name FROM accounts WHERE user_id = ?", userId)
	err = row.Scan(&account.Id, &account.UserId, &account.Password, &account.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repository.ErrResourceNotFound
			return
		}
		log.Printf("[Error]AccountRepository.FindUserId: %v", err)
		err = repository.ErrInternal
		return
	}
	return
}

func (accountRepo *AccountRepository) Create(ctx context.Context,userId string, password string, name string) (account entity.Account, err error) {
	if err = accountRepo.checkConnection(); err != nil {
		return
	}
	result, err := accountRepo.SqlHandler.Conn.Exec("INSERT accounts(user_id, password, name) VALUES (?, ?, ?)", userId, password, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repository.ErrResourceConflict
		}
		log.Printf("[Error]AccountRepository.Create.exec: %v", err)
		err = repository.ErrInternal
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Printf("[Error]AccountRepository.Create.lastInsertId: %v", err)
		err = repository.ErrInternal
		return
	}
	account = entity.Account{
		Id:       int(lastInsertId),
		UserId:   userId,
		Password: password,
		Name:     name,
	}
	return account, nil
}
