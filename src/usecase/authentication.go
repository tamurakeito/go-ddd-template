package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
)

type AuthUsecase interface {
	SignIn(userId string, password string) (account model.Account, token string, err error)
}

type authUsecase struct {
	authRepo    repository.AuthRepository
	authServ    service.AuthService
	encryptServ service.EncryptService
}

func NewAuthUsecase(authRepo repository.AuthRepository, authServ service.AuthService, encryptServ service.EncryptService) AuthUsecase {
	return &authUsecase{
		authRepo:    authRepo,
		authServ:    authServ,
		encryptServ: encryptServ,
	}
}

func (usecase *authUsecase) SignIn(userId string, password string) (account model.Account, token string, err error) {
	// ユーザーをリポジトリから取得
	account, err = usecase.authRepo.FindUserId(userId)
	if err != nil {
		// ユーザーが見つからない場合のエラーハンドリング
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user not found")
			return
		}
		return
	}

	// err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	err = usecase.encryptServ.ComparePassword(account.Password, password)
	if err != nil {
		err = fmt.Errorf("invalid password")
		return
	}

	token, err = usecase.authServ.GenerateToken(userId)
	return
}
