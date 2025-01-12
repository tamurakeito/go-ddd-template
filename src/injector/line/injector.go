package injector

import (
	"crypto/tls"
	"net/http"
	"os"

	repository_line "go-ddd-template/src/domain/repository/line"
	repository_impl_line "go-ddd-template/src/infrastructure/line/repository_impl"
	handler_line "go-ddd-template/src/presentation/handler/line"
	usecase_line "go-ddd-template/src/usecase/line"
)

func InjectHTTPClient() *http.Client {
	// 環境変数から環境を取得（"production" または "development"）
	env := os.Getenv("APP_ENV")
	if env == "production" {
		// 本番環境: 証明書検証を有効化
		return &http.Client{}
	}

	// 開発環境: 証明書検証を無効化
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 証明書検証を無効化
			},
		},
	}
}

func InjectLineMessageRepository() repository_line.LineMessageRepository {
	httpClient := InjectHTTPClient()
	channelToken := "hMQjeZpeqVk8ep/wMfO6kWeW0OKoNtJzsa/I5iJ9eMdtdYjU7UQhDkdi1SLVP4w4IXAgSI2PWSkV3f7TWR41X/BkTqPF/ZDrGPRXtdDIIJ9BnoDkzWMZs33Wh42t9+BLFG8DMvgU0n/UAFkrZgLd7gdB04t89/1O/w1cDnyilFU=" // 環境変数や設定ファイルから取得するのが推奨
	apiEndpoint := "https://api.line.me"
	return repository_impl_line.NewLineMessageRepository(httpClient, channelToken, apiEndpoint)
}

func InjectLineMessageUsecase() usecase_line.LineMessageUsecase {
	lineClient := InjectLineMessageRepository()
	return usecase_line.NewLineMessageUsecase(lineClient)
}

func InjectLineMessageHandler() handler_line.LineMessageHandler {
	return handler_line.NewLinMessageHandler(InjectLineMessageUsecase())
}
