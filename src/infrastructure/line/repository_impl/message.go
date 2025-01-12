package repository_impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	repository_line "go-ddd-template/src/domain/repository/line"
	infrastructure_line "go-ddd-template/src/infrastructure/line"
	"io"
	"net/http"
)

type LineMessageRepository struct {
    httpClient *http.Client
    channelToken string
	apiEndpoint   string
}

func NewLineMessageRepository(httpClient *http.Client, channelToken string,	apiEndpoint string) repository_line.LineMessageRepository {
    lineMessageRepository := LineMessageRepository{
        httpClient: httpClient,
        channelToken: channelToken,
        apiEndpoint: apiEndpoint,
    }
    return &lineMessageRepository
}

func (lineRepo *LineMessageRepository) PushMessage(ctx context.Context, userId string, message string) (err error) {
    reqBody := map[string]interface{}{
        "to": userId,
        "messages": []map[string]string{
            {"type": "text", "text": message},
        },
    }
    reqBytes, _ := json.Marshal(reqBody)

    req, err := http.NewRequest("POST", lineRepo.apiEndpoint+infrastructure_line.PushMessagePath, bytes.NewBuffer(reqBytes))
    if err != nil {
        // ここに想定されるエラーハンドリング
        return err
    }

    req.Header.Set("Authorization", "Bearer "+lineRepo.channelToken)
    req.Header.Set("Content-Type", "application/json")

    resp, err := lineRepo.httpClient.Do(req)
    if err != nil {
        // ここに想定されるエラーハンドリング
        return err
    }
    defer resp.Body.Close()

    // レスポンスボディを読み取る
    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read response body: %w", err)
    }

    // ステータスコード確認
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to send message: status %d, response %s", resp.StatusCode, string(respBody))
    }
    return nil
}
