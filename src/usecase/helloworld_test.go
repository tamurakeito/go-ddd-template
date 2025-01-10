package usecase

import (
	"context"
	"errors"
	mocks "go-ddd-template/mocks/repository"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewHelloWorldUsecase(t *testing.T) {
	type args struct {
		helloRepo repository.HelloRepository
	}
	tests := []struct {
		name string
		args args
		want HelloWorldUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHelloWorldUsecase(tt.args.helloRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelloWorldUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helloWorldUsecase_HelloWorldDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHelloRepo := mocks.NewMockHelloRepository(ctrl)

	type fields struct {
		helloRepo repository.HelloRepository
	}
	type args struct {
		ctx context.Context
		id int
	}
	type test struct {
		name       string
		fields     fields
		args       args
		wantDetail entity.HelloWorld
		wantErr    error
	}
	tests := []test{
		func() test {
            ctx := context.Background()
			id := 1
			hello := entity.Hello{Id: id, Name: "hello, world!", Tag: true}

			mockHelloRepo.EXPECT().
				Find(ctx, id).
				Return(hello, nil).Times(1)

			return test{
				name: "success case",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
					ctx: ctx,
					id: id,
				},
				wantDetail: entity.HelloWorld{
					Id:    id,
					Hello: hello,
				},
				wantErr: nil,
			}
		}(),
		func() test {
            ctx := context.Background()
			id := 3

			mockHelloRepo.EXPECT().
				Find(ctx, id).
				Return(entity.Hello{}, repository.ErrDatabaseUnavailable).Times(1)

			return test{
				name: "database connection failed",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
					ctx: ctx,
					id: id,
				},
				wantDetail: entity.HelloWorld{},
				wantErr:    ErrDatabaseUnavailable,
			}
		}(),
		func() test {
            ctx := context.Background()
			id := 5

			mockHelloRepo.EXPECT().
				Find(ctx, id).
				Return(entity.Hello{}, repository.ErrResourceNotFound).Times(1)

			return test{
				name: "no data case",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
					ctx: ctx,
					id: id,
				},
				wantDetail: entity.HelloWorld{},
				wantErr:    ErrResourceNotFound,
			}
		}(),
		func() test {
            ctx := context.Background()
			id := 999
			err := repository.ErrInternal

			mockHelloRepo.EXPECT().
				Find(ctx, id).
				Return(entity.Hello{}, err).Times(1)

			return test{
				name: "unexpected error case",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
					ctx: ctx,
					id: id,
				},
				wantDetail: entity.HelloWorld{},
				wantErr:    ErrInternal,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &helloWorldUsecase{
				helloRepo: tt.fields.helloRepo,
			}
			gotDetail, err := u.HelloWorldDetail(tt.args.ctx, tt.args.id)

			// エラーが期待と一致しない場合
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("helloWorldUsecase.HelloWorldDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 詳細結果が期待と一致しない場合
			if !reflect.DeepEqual(gotDetail, tt.wantDetail) {
				t.Errorf("helloWorldUsecase.HelloWorldDetail() = %v, want %v", gotDetail, tt.wantDetail)
			}
		})
	}
}
