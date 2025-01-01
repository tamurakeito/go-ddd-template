package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
)

type AccountUsecase interface {
	SignIn(userId string, password string) (account model.Account, token string, err error)
	SignUp(userId string, password string, name string) (account model.Account, token string, err error)
}

type accountUsecase struct {
	accountRepo repository.AccountRepository
	authServ    service.AuthService
	encryptServ service.EncryptService
}

func NewAccountUsecase(accountRepo repository.AccountRepository, authServ service.AuthService, encryptServ service.EncryptService) AccountUsecase {
	return &accountUsecase{
		accountRepo: accountRepo,
		authServ:    authServ,
		encryptServ: encryptServ,
	}
}

func (usecase *accountUsecase) SignIn(userId string, password string) (account model.Account, token string, err error) {
	// ユーザーをリポジトリから取得
	account, err = usecase.accountRepo.FindUserId(userId)
	if err != nil {
		// ユーザーが見つからない場合のエラーハンドリング
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user not found")
			return
		}
		err = fmt.Errorf("failed to retrieve user: %w", err)
		return
	}

	// err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	err = usecase.encryptServ.ComparePassword(account.Password, password)
	if err != nil {
		err = fmt.Errorf("invalid password")
		return
	}

	token, err = usecase.authServ.GenerateToken(userId)
	if err != nil {
		err = fmt.Errorf("failed to generate token: %w", err)
		return
	}
	return
}

func (usecase *accountUsecase) SignUp(userId string, password string, name string) (account model.Account, token string, err error) {
	hashedPassword, err := usecase.encryptServ.HashPassword(password)
	if err != nil {
		err = fmt.Errorf("failed to hash password: %w", err)
		return
	}

	account, err = usecase.accountRepo.Create(userId, hashedPassword, name)
	if err != nil {
		err = fmt.Errorf("failed to create account: %w", err)
		return
	}

	token, err = usecase.authServ.GenerateToken(userId)
	if err != nil {
		err = fmt.Errorf("failed to generate token: %w", err)
		return
	}
	return
}
