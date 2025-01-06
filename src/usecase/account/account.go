package usecase_account

import (
	"database/sql"
	"errors"
	"log"

	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
	"go-ddd-template/src/usecase"
)

type AccountUsecase interface {
	SignIn(userId string, password string) (account entity.Account, token string, err error)
	SignUp(userId string, password string, name string) (account entity.Account, token string, err error)
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

func (u *accountUsecase) SignIn(userId string, password string) (account entity.Account, token string, err error) {
	// ユーザーをリポジトリから取得
	account, err = u.accountRepo.FindUserId(userId)
	if err != nil {
		// ユーザーが見つからない場合のエラーハンドリング
		if errors.Is(err, sql.ErrNoRows) {
			err = usecase.ErrUserNotFound
			return
		}
		err = usecase.ErrFailedToRetrieveUser
		log.Printf("[SignIn] Error retrieving user: %v", err)
		return
	}

	// err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	err = u.encryptServ.ComparePassword(account.Password, password)
	if err != nil {
		err = usecase.ErrInvalidPassword
		log.Printf("[SignIn] Password verification failed: %v", err)
		return
	}

	token, err = u.authServ.GenerateToken(userId)
	if err != nil {
		err = usecase.ErrFailedToGenerateToken
		log.Printf("[SignIn] Token generation failed: %v", err)
		return
	}
	return
}

func (u *accountUsecase) SignUp(userId string, password string, name string) (account entity.Account, token string, err error) {
	hashedPassword, err := u.encryptServ.HashPassword(password)
	if err != nil {
		err = usecase.ErrFailedToHashPassword
		log.Printf("[SignUp] Password hashing failed: %v", err)
		return
	}

	account, err = u.accountRepo.Create(userId, hashedPassword, name)
	if err != nil {
		err = usecase.ErrFailedToCreateAccount
		log.Printf("[SignUp] Account creation failed: %v", err)
		return
	}

	token, err = u.authServ.GenerateToken(userId)
	if err != nil {
		err = usecase.ErrFailedToGenerateToken
		log.Printf("[SignUp] Token generation failed: %v", err)
		return
	}
	return
}
