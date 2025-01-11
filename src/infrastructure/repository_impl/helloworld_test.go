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
)

func TestNewHelloRepository(t *testing.T) {
	type args struct {
		sqlHandler infrastructure.SqlHandler
	}
	tests := []struct {
		name string
		args args
		want repository.HelloRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHelloRepository(tt.args.sqlHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelloRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelloRepository_Find(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to create sqlmock: %v", err)
    }
    defer db.Close()

	sqlHandler := infrastructure.SqlHandler{Conn: db}
	type fields struct {
		SqlHandler infrastructure.SqlHandler
	}
	type args struct {
		ctx context.Context
		id  int
	}
	type test struct {
        name    string
		fields     fields
        args    args
        wantHello    entity.Hello
        wantErr error
    }

	tests := []test{
		func() test {
			ctx := context.Background()
			rows := sqlmock.NewRows([]string{"id", "name", "tag"}).
				AddRow(1, "TestName", true)
			mock.ExpectQuery("SELECT id, name, tag FROM hello_world WHERE id = ?").
				WithArgs(1).
				WillReturnRows(rows)

			return test{
				name: "success case",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					id: 1,
				},
				wantHello: entity.Hello{
					Id:   1,
					Name: "TestName",
					Tag:  true,
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
					id: 1,
				},
				wantHello: entity.Hello{},
				wantErr: repository.ErrDatabaseUnavailable,
	
			}
		}(),
		func() test {
			ctx := context.Background()
			rows := sqlmock.NewRows([]string{"id", "name", "tag"}).
				AddRow(1, "TestName", true)
			mock.ExpectQuery("SELECT id, name, tag FROM hello_world WHERE id = ?").
					WithArgs(1).
					WillDelayFor(4 * time.Second).
					WillReturnRows(rows)

			return test{
				name: "database timeout",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					id: 1,
				},
				wantHello: entity.Hello{},
				wantErr: repository.ErrDatabaseUnavailable,
	
			}
		}(),
		func() test {
			ctx := context.Background()
			mock.ExpectQuery("SELECT id, name, tag FROM hello_world WHERE id = ?").
				WithArgs(999).
				WillReturnError(sql.ErrNoRows)


			return test{
				name: "record not found",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					id: 999,
				},
				wantHello: entity.Hello{},
				wantErr: repository.ErrResourceNotFound,
	
			}
		}(),
		func() test {
			ctx := context.Background()
			mock.ExpectQuery("SELECT id, name, tag FROM hello_world WHERE id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrConnDone) 

			return test{
				name: "unexpected error",
				fields: fields{
					SqlHandler: sqlHandler,
				},
				args: args{
					ctx: ctx,
					id:  1,
				},
				wantHello: entity.Hello{},
				wantErr:   repository.ErrInternal,
			}
		}(),
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helloRepo := &HelloRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			gotHello, err := helloRepo.Find(tt.args.ctx, tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("HelloRepository.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHello, tt.wantHello) {
				t.Errorf("HelloRepository.Find() = %v, want %v", gotHello, tt.wantHello)
			}
		})
	}
}
