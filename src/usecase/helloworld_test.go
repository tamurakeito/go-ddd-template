package usecase

import (
	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
	"reflect"
	"testing"
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
	type fields struct {
		helloRepo repository.HelloRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantDetail model.HelloWorld
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &helloWorldUsecase{
				helloRepo: tt.fields.helloRepo,
			}
			gotDetail, err := usecase.HelloWorldDetail(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("helloWorldUsecase.HelloWorldDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDetail, tt.wantDetail) {
				t.Errorf("helloWorldUsecase.HelloWorldDetail() = %v, want %v", gotDetail, tt.wantDetail)
			}
		})
	}
}
