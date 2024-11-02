package usecase

import (
	"fmt"
	"log"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
)

type AuthUsecase interface {
	SignIn(userId string, password string) (account model.Account, token string, err error)
}

type authUsecase struct {
	authRepo       repository.AuthRepository
	tokenGenerator service.TokenGenerator
}

func NewAuthUsecase(authRepo repository.AuthRepository, tokenGen service.TokenGenerator) AuthUsecase {
	return &authUsecase{
		authRepo:       authRepo,
		tokenGenerator: tokenGen,
	}
}

func (usecase *authUsecase) SignIn(userId string, password string) (account model.Account, token string, err error) {
	// ユーザーをリポジトリから取得
	account, err = usecase.authRepo.FindUserId(userId)
	if err != nil {
		log.Fatal(err)
		return
	}

	// パスワードの検証 (例)
	if account.Password != password {
		err = fmt.Errorf("invalid password")
		return
	}

	// トークンの生成
	token, err = usecase.tokenGenerator.GenerateToken(userId)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
