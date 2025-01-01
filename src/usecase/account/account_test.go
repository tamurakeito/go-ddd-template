package usecase_account

import (
	"database/sql"
	"fmt"
	"go-ddd-template/mocks"
	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
	"go-ddd-template/src/usecase"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewAccountUsecase(t *testing.T) {
	type args struct {
		accountRepo repository.AccountRepository
		authServ    service.AuthService
		encryptServ service.EncryptService
	}
	tests := []struct {
		name string
		args args
		want AccountUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountUsecase(tt.args.accountRepo, tt.args.authServ, tt.args.encryptServ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
	mockAuthServ := mocks.NewMockAuthService(ctrl)
	mockEncryptServ := mocks.NewMockEncryptService(ctrl)

	type fields struct {
		accountRepo repository.AccountRepository
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

			mockAccountRepo.EXPECT().
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
					accountRepo: mockAccountRepo,
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

			mockAccountRepo.EXPECT().
				FindUserId(unknownUser).
				Return(nilAccount, sql.ErrNoRows)

			return test{
				name: "User not found",
				fields: fields{
					accountRepo: mockAccountRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   unknownUser,
					password: anyPaasword,
				},
				wantAccount: nilAccount,
				wantToken:   "",
				wantErr:     usecase.ErrUserNotFound,
			}
		}(),
		func() test {
			userId := "validUser"
			hashedPassword := "hashedPassword"
			wrongPassword := "wrongPassword"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: "Valid User"}
			err := fmt.Errorf("invalid password")

			mockAccountRepo.EXPECT().
				FindUserId(userId).
				Return(account, nil)
			mockEncryptServ.EXPECT().
				ComparePassword(hashedPassword, wrongPassword).
				Return(err)

			return test{
				name: "Invalid password",
				fields: fields{
					accountRepo: mockAccountRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: wrongPassword,
				},
				wantAccount: account,
				wantToken:   "",
				wantErr:     usecase.ErrInvalidPassword,
			}
		}(),
		func() test {
			userId := "validUser"
			password := "validPassword"
			hashedPassword := "hashedPassword"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: "Valid User"}
			err := fmt.Errorf("failed to generate token")

			mockAccountRepo.EXPECT().
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
					accountRepo: mockAccountRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
				},
				wantAccount: account,
				wantToken:   "",
				wantErr:     usecase.ErrFailedToGenerateToken,
			}
		}(),
		func() test {
			userId := "anyUser"
			nilAccount := model.Account{}
			err := fmt.Errorf("unexpected database error")

			mockAccountRepo.EXPECT().
				FindUserId(userId).
				Return(nilAccount, err)

			return test{
				name: "Unexpected error",
				fields: fields{
					accountRepo: mockAccountRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: "anyPassword",
				},
				wantAccount: nilAccount,
				wantToken:   "",
				wantErr:     usecase.ErrFailedToRetrieveUser,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &accountUsecase{
				accountRepo: tt.fields.accountRepo,
				authServ:    tt.fields.authServ,
				encryptServ: tt.fields.encryptServ,
			}
			gotAccount, gotToken, err := usecase.SignIn(tt.args.userId, tt.args.password)
			if (err != nil) && (tt.wantErr != nil) {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("accountUsecase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if (err == nil) != (tt.wantErr == nil) {
				t.Errorf("accountUsecase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("accountUsecase.SignIn() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}
			if gotToken != tt.wantToken {
				t.Errorf("accountUsecase.SignIn() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func Test_accountUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
	mockAuthServ := mocks.NewMockAuthService(ctrl)
	mockEncryptServ := mocks.NewMockEncryptService(ctrl)

	type fields struct {
		accountRepo repository.AccountRepository
		authServ    service.AuthService
		encryptServ service.EncryptService
	}
	type args struct {
		userId   string
		password string
		name     string
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
			name := "Vaild User"
			token := "validToken"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: name}

			mockAccountRepo.EXPECT().
				Create(userId, hashedPassword, name).
				Return(account, nil).Times(1)
			mockAuthServ.EXPECT().
				GenerateToken(userId).
				Return(token, nil).Times(1)
			mockEncryptServ.EXPECT().
				HashPassword(password).
				Return(hashedPassword, nil).Times(1)

			return test{
				name: "success case",
				fields: fields{
					accountRepo: mockAccountRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
					name:     name,
				},
				wantAccount: account,
				wantToken:   token,
				wantErr:     nil,
			}
		}(),
		func() test {
			userId := "validUser"
			password := "invalidPassword"
			hashedPassword := ""
			name := "Valid User"
			err := fmt.Errorf("failed to hash password")

			mockEncryptServ.EXPECT().
				HashPassword(password).
				Return(hashedPassword, err).Times(1)

			return test{
				name: "hash password error",
				fields: fields{
					accountRepo: mockAccountRepo,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
					name:     name,
				},
				wantAccount: model.Account{},
				wantToken:   "",
				wantErr:     usecase.ErrFailedToHashPassword,
			}
		}(),
		func() test {
			userId := "validUser"
			password := "validPassword"
			hashedPassword := "hashedPassword"
			name := "Valid User"
			account := model.Account{}
			err := fmt.Errorf("failed to create account")

			mockEncryptServ.EXPECT().
				HashPassword(password).
				Return(hashedPassword, nil).Times(1)
			mockAccountRepo.EXPECT().
				Create(userId, hashedPassword, name).
				Return(account, err).Times(1)

			return test{
				name: "create account error",
				fields: fields{
					accountRepo: mockAccountRepo,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
					name:     name,
				},
				wantAccount: model.Account{},
				wantToken:   "",
				wantErr:     usecase.ErrFailedToCreateAccount,
			}
		}(),
		func() test {
			userId := "validUser"
			password := "validPassword"
			hashedPassword := "hashedPassword"
			name := "Valid User"
			account := model.Account{Id: 1, UserId: userId, Password: hashedPassword, Name: name}
			token := ""
			err := fmt.Errorf("failed to generate token")

			mockEncryptServ.EXPECT().
				HashPassword(password).
				Return(hashedPassword, nil).Times(1)
			mockAccountRepo.EXPECT().
				Create(userId, hashedPassword, name).
				Return(account, nil).Times(1)
			mockAuthServ.EXPECT().
				GenerateToken(userId).
				Return(token, err).Times(1)

			return test{
				name: "generate token error",
				fields: fields{
					accountRepo: mockAccountRepo,
					authServ:    mockAuthServ,
					encryptServ: mockEncryptServ,
				},
				args: args{
					userId:   userId,
					password: password,
					name:     name,
				},
				wantAccount: account,
				wantToken:   "",
				wantErr:     usecase.ErrFailedToGenerateToken,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &accountUsecase{
				accountRepo: tt.fields.accountRepo,
				authServ:    tt.fields.authServ,
				encryptServ: tt.fields.encryptServ,
			}
			gotAccount, gotToken, err := usecase.SignUp(tt.args.userId, tt.args.password, tt.args.name)
			if (err != nil) && (tt.wantErr != nil) {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("accountUsecase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if (err == nil) != (tt.wantErr == nil) {
				t.Errorf("accountUsecase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("accountUsecase.SignUp() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}
			if gotToken != tt.wantToken {
				t.Errorf("accountUsecase.SignIn() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
