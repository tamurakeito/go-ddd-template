package repository_impl

import (
	"context"
	"errors"
	repository_line "go-ddd-template/src/domain/repository/line"
	infrastructure_line "go-ddd-template/src/infrastructure/line"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewLineMessageRepository(t *testing.T) {
	type args struct {
		httpClient   *http.Client
		channelToken string
		apiEndpoint string
	}
	tests := []struct {
		name string
		args args
		want repository_line.LineMessageRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLineMessageRepository(tt.args.httpClient, tt.args.channelToken, tt.args.apiEndpoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineMessageRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineMessageRepository_PushMessage(t *testing.T) {
	type fields struct {
		httpClient   *http.Client
		channelToken string
		apiEndpoint string
	}
	type args struct {
		ctx     context.Context
		userID  string
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
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != infrastructure_line.PushMessagePath {
					t.Errorf("unexpected URL path: %s", r.URL.Path)
				}
				w.WriteHeader(http.StatusOK)
			}))
			mockHTTPClient := server.Client()

			return test{
				name: "success case",
				fields: fields{
					httpClient:   mockHTTPClient,
					channelToken: "TEST_CHANNEL_TOKEN",
					apiEndpoint: server.URL,
				},
				args: args{
					ctx:     context.Background(),
					userID:  "USER_ID",
					message: "Hello, World!",
				},
				wantErr: nil,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lineRepo := &LineMessageRepository{
				httpClient:   tt.fields.httpClient,
				channelToken: tt.fields.channelToken,
				apiEndpoint: tt.fields.apiEndpoint,
			}
			if err := lineRepo.PushMessage(tt.args.ctx, tt.args.userID, tt.args.message); !errors.Is(err, tt.wantErr)  {
				t.Errorf("LineMessageRepository.PushMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
