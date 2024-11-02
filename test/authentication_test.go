package test

import (
	"reflect"
	"testing"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/usecase"
	"go-ddd-template/test/mocks"

	"github.com/golang/mock/gomock"
)

func TestAuthUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックのセットアップ
	mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
	mockTokenGen := mocks.NewMockTokenGenerator(ctrl)
	userId := "user"
	password := "password"
	expectedAccount := model.Account{Id: 0, UserId: "user", Password: "password", Name: "Test User"}
	expectedToken := "mockToken"

	// モックの動作を定義
	mockAuthRepo.EXPECT().FindUserId(userId).Return(expectedAccount, nil).Times(1)
	mockTokenGen.EXPECT().GenerateToken(userId).Return(expectedToken, nil).Times(1)

	// テスト対象のユースケースを作成
	authUsecase := usecase.NewAuthUsecase(mockAuthRepo, mockTokenGen)

	// SignInメソッドを呼び出し
	account, token, err := authUsecase.SignIn(userId, password)

	// テストの検証
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(account, expectedAccount) {
		t.Errorf("expected account: %v, got: %v", expectedAccount, account)
	}
	if token != expectedToken {
		t.Errorf("expected token: %v, got: %v", expectedToken, token)
	}
}
