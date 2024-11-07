package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHelloWorldUsecase_HelloWorldDetail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}
