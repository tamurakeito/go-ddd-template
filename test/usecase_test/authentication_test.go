package usecase_test

import (
	"database/sql"
	"errors"
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

	// 共通のモックセットアップ
	mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
	mockAuthServ := mocks.NewMockAuthService(ctrl)
	userId := "user"
	password := "password"
	expectedAccount := model.Account{Id: 0, UserId: "user", Password: "password", Name: "Test User"}
	expectedToken := "mockToken"

	// --- 正常なケースのテスト ---
	mockAuthRepo.EXPECT().FindUserId(userId).Return(expectedAccount, nil).Times(1)
	mockAuthServ.EXPECT().GenerateToken(userId).Return(expectedToken, nil).Times(1)

	// ユースケースのインスタンスを作成
	authUsecase := usecase.NewAuthUsecase(mockAuthRepo, mockAuthServ)

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

	// --- パスワードが誤っている場合のテスト ---
	mockAuthRepo.EXPECT().FindUserId(userId).Return(expectedAccount, nil).Times(1) // ユーザーは見つかる
	expectedError := "invalid password"
	_, _, err = authUsecase.SignIn(userId, "wrongpassword")
	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error: %v, got: %v", expectedError, err)
	}

	// --- ユーザーが存在しない場合のテスト ---
	mockAuthRepo.EXPECT().FindUserId(userId).Return(model.Account{}, sql.ErrNoRows).Times(1)
	_, _, err = authUsecase.SignIn(userId, password)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("expected error: user not found, got: %v", err)
	}

	// --- 予期しないエラーが発生した場合のテスト ---
	unexpectedError := errors.New("database connection error")
	mockAuthRepo.EXPECT().FindUserId(userId).Return(model.Account{}, unexpectedError).Times(1)
	_, _, err = authUsecase.SignIn(userId, password)
	if err == nil || err.Error() != unexpectedError.Error() {
		t.Errorf("expected error: %v, got: %v", unexpectedError.Error(), err)
	}
}
