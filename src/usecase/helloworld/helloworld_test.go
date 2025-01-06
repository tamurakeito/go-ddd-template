package usecase_helloworld

import (
	"database/sql"
	"fmt"
	"go-ddd-template/mocks"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/usecase"
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
			id := 1
			hello := entity.Hello{Id: id, Name: "hello, world!", Tag: true}

			mockHelloRepo.EXPECT().
				Find(id).
				Return(hello, nil).Times(1)

			return test{
				name: "success case",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
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
			id := 5

			mockHelloRepo.EXPECT().
				Find(id).
				Return(entity.Hello{}, sql.ErrNoRows).Times(1)

			return test{
				name: "no data case",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
					id: id,
				},
				wantDetail: entity.HelloWorld{},
				wantErr:    usecase.ErrResultsNotFound,
			}
		}(),
		func() test {
			id := 999
			err := fmt.Errorf("unexpected error")

			mockHelloRepo.EXPECT().
				Find(id).
				Return(entity.Hello{}, err).Times(1)

			return test{
				name: "unexpected error case",
				fields: fields{
					helloRepo: mockHelloRepo,
				},
				args: args{
					id: id,
				},
				wantDetail: entity.HelloWorld{},
				wantErr:    usecase.ErrFailedToRetrieveData,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &helloWorldUsecase{
				helloRepo: tt.fields.helloRepo,
			}
			gotDetail, err := usecase.HelloWorldDetail(tt.args.id)
			if (err != nil) && (tt.wantErr != nil) {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("helloWorldUsecase.HelloWorldDetail() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if (err == nil) != (tt.wantErr == nil) {
				t.Errorf("helloWorldUsecase.HelloWorldDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDetail, tt.wantDetail) {
				t.Errorf("helloWorldUsecase.HelloWorldDetail() = %v, want %v", gotDetail, tt.wantDetail)
			}
		})
	}
}
