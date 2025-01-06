package usecase

import (
	"errors"
	"log"

	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
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
		if errors.Is(err, repository.ErrResourceNotFound) {
			err = ErrResourceNotFound
			return
		}
		err = ErrInternal
		return
	}

	err = u.encryptServ.ComparePassword(account.Password, password)
	if err != nil {
		if errors.Is(err, repository.ErrResourceNotFound) {
			log.Printf("[Error]AccountUsecase.ComparePassword: invalid password")
			err = ErrResourceNotFound // セキュリティの観点からユーザーが存在しない場合と同じエラーを返す
			return
		}
		log.Printf("[Error]AccountUsecase.ComparePassword: %v", err)
		err = ErrInternal
		return
	}

	token, err = u.authServ.GenerateToken(userId)
	if err != nil {
		log.Printf("[Error]AccountUsecase.GenerateToken: %v", err)
		err = ErrInternal
		return
	}
	return
}

func (u *accountUsecase) SignUp(userId string, password string, name string) (account entity.Account, token string, err error) {
	hashedPassword, err := u.encryptServ.HashPassword(password)
	if err != nil {
		log.Printf("[Error]AccountUsecase.HashPassword: %v", err)
		err = ErrInternal
		return
	}

	account, err = u.accountRepo.Create(userId, hashedPassword, name)
	if err != nil {
		if errors.Is(err, repository.ErrResourceConflict){
			err = ErrResourceConflict
			return
		}
		log.Printf("[Error]AccountUsecase.HashPassword: %v", err)
		err = ErrInternal
		return
	}

	token, err = u.authServ.GenerateToken(userId)
	if err != nil {
		log.Printf("[Error]AccountUsecase.GenerateToken: %v", err)
		err = ErrInternal
		return
	}
	return
}
