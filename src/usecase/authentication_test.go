package usecase

import (
	"database/sql"
	"fmt"
	"go-ddd-template/mocks"
	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewAuthUsecase(t *testing.T) {
	type args struct {
		authRepo    repository.AuthRepository
		authServ    service.AuthService
		encryptServ service.EncryptService
	}
	tests := []struct {
		name string
		args args
		want AuthUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecase(tt.args.authRepo, tt.args.authServ, tt.args.encryptServ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
	mockAuthServ := mocks.NewMockAuthService(ctrl)
	mockEncryptServ := mocks.NewMockEncryptService(ctrl)

	type fields struct {
		authRepo    repository.AuthRepository
		authServ    service.AuthService
		encryptServ service.EncryptService
	}
	type args struct {
		userId   string
		password string
	}

	type test struct {
		name        string
		fields      fields
		args        args
		wantAccount model.Account
		wantToken   string
		wantErr     error
	}

	tests := []test{
		func() test {
			userId := "validUser"
			password := "validPassword"
			hashedPassword := "hashedPassword"
			token := "validToken"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: "Vaild User"}

			mockAuthRepo.EXPECT().
				FindUserId(userId).
				Return(account, nil).Times(1)
			mockAuthServ.EXPECT().
				GenerateToken(userId).
				Return(token, nil).Times(1)
			mockEncryptServ.EXPECT().
				ComparePassword(hashedPassword, password).
				Return(nil).Times(1)

			return test{
				name: "valid user credentials",
				fields: fields{
					authRepo:    mockAuthRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
				},
				wantAccount: account,
				wantToken:   token,
				wantErr:     nil,
			}
		}(),
		func() test {
			unknownUser := "unknownUser"
			anyPaasword := "anyPassword"
			nilAccount := model.Account{}

			mockAuthRepo.EXPECT().
				FindUserId(unknownUser).
				Return(nilAccount, sql.ErrNoRows)

			return test{
				name: "User not found",
				fields: fields{
					authRepo:    mockAuthRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   unknownUser,
					password: anyPaasword,
				},
				wantAccount: nilAccount,
				wantToken:   "",
				wantErr:     fmt.Errorf("user not found"),
			}
		}(),
		func() test {
			userId := "validUser"
			hashedPassword := "hashedPassword"
			wrongPassword := "wrongPassword"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: "Valid User"}
			err := fmt.Errorf("invalid password")

			mockAuthRepo.EXPECT().
				FindUserId(userId).
				Return(account, nil)
			mockEncryptServ.EXPECT().
				ComparePassword(hashedPassword, wrongPassword).
				Return(err)

			return test{
				name: "Invalid password",
				fields: fields{
					authRepo:    mockAuthRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: wrongPassword,
				},
				wantAccount: account,
				wantToken:   "",
				wantErr:     err,
			}
		}(),
		func() test {
			userId := "validUser"
			password := "validPassword"
			hashedPassword := "hashedPassword"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: "Valid User"}
			err := fmt.Errorf("token generation failed")

			mockAuthRepo.EXPECT().
				FindUserId(userId).
				Return(account, nil)
			mockEncryptServ.EXPECT().
				ComparePassword(hashedPassword, password).
				Return(nil)
			mockAuthServ.EXPECT().
				GenerateToken(userId).
				Return("", err)

			return test{
				name: "Token generation error",
				fields: fields{
					authRepo:    mockAuthRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
				},
				wantAccount: account,
				wantToken:   "",
				wantErr:     err,
			}
		}(),
		func() test {
			userId := "anyUser"
			nilAccount := model.Account{}
			err := fmt.Errorf("unexpected database error")

			mockAuthRepo.EXPECT().
				FindUserId(userId).
				Return(nilAccount, err)

			return test{
				name: "Unexpected error",
				fields: fields{
					authRepo:    mockAuthRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: "anyPassword",
				},
				wantAccount: nilAccount,
				wantToken:   "",
				wantErr:     err,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &authUsecase{
				authRepo:    tt.fields.authRepo,
				authServ:    tt.fields.authServ,
				encryptServ: tt.fields.encryptServ,
			}
			gotAccount, gotToken, err := usecase.SignIn(tt.args.userId, tt.args.password)
			if (err != nil) && (tt.wantErr != nil) {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("authUsecase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if (err == nil) != (tt.wantErr == nil) {
				t.Errorf("authUsecase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("authUsecase.SignIn() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}
			if gotToken != tt.wantToken {
				t.Errorf("authUsecase.SignIn() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
