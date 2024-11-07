package usecase

import (
	"database/sql"
	"fmt"
	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/service"
	"go-ddd-template/test/mocks"
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

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAccount model.Account
		wantToken   string
		wantErr     error
		setupMocks  func()
	}{
		{
			name: "valid user credentials",
			fields: fields{
				authRepo:    mockAuthRepo,
				authServ:    mockAuthServ,
				encryptServ: mockEncryptServ,
			},
			args: args{
				userId:   "validUser",
				password: "validPassword",
			},
			wantAccount: model.Account{Id: 1, UserId: "validUser", Password: "hashedPassword", Name: "Valid User"},
			wantToken:   "validToken",
			wantErr:     nil,
			setupMocks: func() {
				mockAuthRepo.EXPECT().
					FindUserId("validUser").
					Return(model.Account{Id: 1, UserId: "validUser", Password: "hashedPassword", Name: "Valid User"}, nil).Times(1)
				mockAuthServ.EXPECT().
					GenerateToken("validUser").
					Return("validToken", nil).Times(1)
				mockEncryptServ.EXPECT().
					ComparePassword("hashedPassword", "validPassword").
					Return(nil).Times(1)
			},
		},
		{
			name: "User not found",
			fields: fields{
				authRepo:    mockAuthRepo,
				authServ:    mockAuthServ,
				encryptServ: mockEncryptServ,
			},
			args: args{
				userId:   "unknownUser",
				password: "anyPassword",
			},
			wantAccount: model.Account{},
			wantToken:   "",
			wantErr:     fmt.Errorf("user not found"),
			setupMocks: func() {
				mockAuthRepo.EXPECT().
					FindUserId("unknownUser").
					Return(model.Account{}, sql.ErrNoRows)
			},
		},
		{
			name: "Invalid password",
			fields: fields{
				authRepo:    mockAuthRepo,
				authServ:    mockAuthServ,
				encryptServ: mockEncryptServ,
			},
			args: args{
				userId:   "validUser",
				password: "wrongPassword",
			},
			wantAccount: model.Account{Id: 1, UserId: "validUser", Password: "hashedPassword", Name: "Valid User"},
			wantToken:   "",
			wantErr:     fmt.Errorf("invalid password"),
			setupMocks: func() {
				mockAuthRepo.EXPECT().
					FindUserId("validUser").
					Return(model.Account{Id: 1, UserId: "validUser", Password: "hashedPassword", Name: "Valid User"}, nil)
				mockEncryptServ.EXPECT().
					ComparePassword("hashedPassword", "wrongPassword").
					Return(fmt.Errorf("invalid password"))
			},
		}, {
			name: "Token generation error",
			fields: fields{
				authRepo:    mockAuthRepo,
				authServ:    mockAuthServ,
				encryptServ: mockEncryptServ,
			},
			args: args{
				userId:   "validUser",
				password: "validPassword",
			},
			wantAccount: model.Account{Id: 1, UserId: "validUser", Password: "hashedPassword", Name: "Valid User"},
			wantToken:   "",
			wantErr:     fmt.Errorf("token generation failed"),
			setupMocks: func() {
				mockAuthRepo.EXPECT().
					FindUserId("validUser").
					Return(model.Account{Id: 1, UserId: "validUser", Password: "hashedPassword", Name: "Valid User"}, nil)
				mockEncryptServ.EXPECT().
					ComparePassword("hashedPassword", "validPassword").
					Return(nil)
				mockAuthServ.EXPECT().
					GenerateToken("validUser").
					Return("", fmt.Errorf("token generation failed"))
			},
		},
		{
			name: "Unexpected error",
			fields: fields{
				authRepo:    mockAuthRepo,
				authServ:    mockAuthServ,
				encryptServ: mockEncryptServ,
			},
			args: args{
				userId:   "anyUser",
				password: "anyPassword",
			},
			wantAccount: model.Account{},
			wantToken:   "",
			wantErr:     fmt.Errorf("unexpected database error"),
			setupMocks: func() {
				mockAuthRepo.EXPECT().
					FindUserId("anyUser").
					Return(model.Account{}, fmt.Errorf("unexpected database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

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
