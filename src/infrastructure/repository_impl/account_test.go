package repository_impl

import (
	"context"
	"database/sql"
	"errors"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
)

func TestNewAccountRepository(t *testing.T) {
	type args struct {
		sqlHandler infrastructure.SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.AccountRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountRepository(tt.args.sqlHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_FindUserId(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to create sqlmock: %v", err)
    }
    defer db.Close()

	sqlHandler := infrastructure.SqlHandler{Conn: db}
	query := "SELECT id, user_id, password, name FROM accounts WHERE user_id = ?"
	type fields struct {
		SqlHandler infrastructure.SqlHandler
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	type test struct {
		name        string
		fields      fields
		args        args
		wantAccount entity.Account
		wantErr     error
	}

	tests := [] test {
		func() test {
			ctx := context.Background()
			id := 1
			userId := "user"
			password := "password"
			name := "valid user"
			rows := sqlmock.NewRows([]string{"id", "user_id", "password","name"}).
				AddRow(id, userId, password, name)
			mock.ExpectQuery(query).
				WithArgs(userId).
				WillReturnRows(rows)

			return test{
				name: "success case",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
				},
				wantAccount: entity.Account{
					Id: id,
					UserId: userId,
					Password: password,
					Name: name,
				},
				wantErr: nil,
			}
		}(),
		func() test {
			ctx := context.Background()

			return test{
				name: "database connection is nil",
				fields: fields{
					SqlHandler: infrastructure.SqlHandler{Conn: nil},
				},
				args: args{
					ctx: ctx,
					userId: "user",
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrDatabaseUnavailable,
			}
		}(),
		func() test {
			ctx := context.Background()
			userId := "user"
			rows := sqlmock.NewRows([]string{"id", "name", "tag"}).
				AddRow(1, "TestName", true)
				mock.ExpectQuery(query).
					WithArgs(userId).
					WillDelayFor(4 * time.Second).
					WillReturnRows(rows)
					
			return test{
				name: "database timeout",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrDatabaseUnavailable,
			}
		}(),
		func() test {
			ctx := context.Background()
			userId := "invalidUser"
			mock.ExpectQuery(query).
				WithArgs(userId).
				WillReturnError(sql.ErrNoRows)

			return test{
				name: "record not found",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrResourceNotFound,
			}
		}(),
		func() test {
			ctx := context.Background()
			userId := "user"
			mock.ExpectQuery(query).
				WithArgs(userId).
				WillReturnError(sql.ErrConnDone)

			return test{
				name: "unexpected error",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrInternal,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountRepo := &AccountRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			gotAccount, err := accountRepo.FindUserId(tt.args.ctx, tt.args.userId)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("AccountRepository.FindUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("AccountRepository.FindUserId() = %v, want %v", gotAccount, tt.wantAccount)
			}
		})
	}
}

func TestAccountRepository_Create(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to create sqlmock: %v", err)
    }
    defer db.Close()

	sqlHandler := infrastructure.SqlHandler{Conn: db}
	query := "INSERT accounts\\(user_id, password, name\\) VALUES \\(\\?, \\?, \\?\\)"
	type fields struct {
		SqlHandler infrastructure.SqlHandler
	}
	type args struct {
		ctx      context.Context
		userId   string
		password string
		name     string
	}
	type test struct {
		name        string
		fields      fields
		args        args
		wantAccount entity.Account
		wantErr     error
	}

	tests := []test{
		func() test {
			ctx := context.Background()
			id := 1
			userId := "user"
			password := "password"
			name := "valid user"
			mock.ExpectExec(query).
				WithArgs(userId, password, name).
				WillReturnResult(sqlmock.NewResult(1, 1))

			return test{
				name: "success case",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId   :userId,
					password :password,
					name     :name,
				},
				wantAccount: entity.Account{
					Id: id,
					UserId: userId,
					Password: password,
					Name: name,
				},
				wantErr: nil,
			}
		}(),
		func() test {
			ctx := context.Background()

			return test{
				name: "database connection is nil",
				fields: fields{
					SqlHandler: infrastructure.SqlHandler{Conn: nil},
				},
				args: args{
					ctx: ctx,
					userId: "user",
					password: "password",
					name: "valid user",
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrDatabaseUnavailable,
	
			}
		}(),
		func() test {
			ctx := context.Background()
			userId := "user"
			password := "password"
			name := "valid user"
			mock.ExpectExec(query).
				WithArgs(userId, password, name).
				WillDelayFor(4 * time.Second).
				WillReturnResult(sqlmock.NewResult(1, 1))
					
			return test{
				name: "database timeout",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
					password: password,
					name: name,
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrDatabaseUnavailable,
	
			}
		}(),
		func() test {
			ctx := context.Background()
			userId := "duplicatedUser"
			password := "password"
			name := "duplicated user"
			mock.ExpectExec(query).
				WithArgs(userId, password, name).
				WillReturnError(&mysql.MySQLError{
					Number:  1062,
					Message: "Duplicate entry 'test_user' for key 'user_id'",
				})
		

			return test{
				name: "unique constraint violation",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
					password: password,
					name: name,
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrResourceConflict,
			}
		}(),
		func() test {
			ctx := context.Background()
			userId := "user"
			password := "password"
			name := "valid user"
			mock.ExpectExec(query).
				WithArgs(userId, password, name).
				WillReturnError(sql.ErrConnDone)

			return test{
				name: "unexpected error",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					userId: userId,
					password: password,
					name: name,
				},
				wantAccount: entity.Account{},
				wantErr: repository.ErrInternal,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountRepo := &AccountRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			gotAccount, err := accountRepo.Create(tt.args.ctx, tt.args.userId, tt.args.password, tt.args.name)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("AccountRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("AccountRepository.Create() = %v, want %v", gotAccount, tt.wantAccount)
			}
		})
	}
}
