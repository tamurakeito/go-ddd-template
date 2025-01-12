package usecase

import (
	"context"
	"errors"
	mocks "go-ddd-template/mocks/repository/line"
	repository_line "go-ddd-template/src/domain/repository/line"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewLineMessageUsecase(t *testing.T) {
	type args struct {
		lineRepo repository_line.LineMessageRepository
	}
	tests := []struct {
		name string
		args args
		want LineMessageUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLineMessageUsecase(tt.args.lineRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineMessageUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lineMessageUsecase_NotifyUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLineRepo := mocks.NewMockLineMessageRepository(ctrl)

	type fields struct {
		lineRepo repository_line.LineMessageRepository
	}
	type args struct {
		ctx     context.Context
		userId  string
		message string
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}

	tests := []test{
		func() test {
            ctx := context.Background()
			userId := "user"
			message := "Hello, world!"

			mockLineRepo.EXPECT().
				PushMessage(ctx, userId, message).
				Return(nil).Times(1)
			
			return test{
				name: "success case",
				fields: fields{
					lineRepo: mockLineRepo,
				},
				args: args{
					ctx: ctx,
					userId: userId,
					message: message,
				},
				wantErr: nil,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &lineMessageUsecase{
				lineRepo: tt.fields.lineRepo,
			}
			if err := u.NotifyUser(tt.args.ctx, tt.args.userId, tt.args.message);  !errors.Is(err, tt.wantErr){
				t.Errorf("lineMessageUsecase.NotifyUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
