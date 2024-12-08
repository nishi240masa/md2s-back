package qitta

import (
	"encoding/json"
	"fmt"
	"io"
	"md2s/models"
	"net/http"
	"os"
	"strings"
)

func GetQiitaAccessToken(code string) (string, error) {

	var client_id = os.Getenv("QIITA_CLIENT_ID")
	var client_secret = os.Getenv("QIITA_CLIENT_SECRET")

	api := "https://qiita.com/api/v2/access_tokens"



	res, err := http.Post(api, "application/json", strings.NewReader(`{"client_id":"`+client_id+`","client_secret":"`+client_secret+`","code":"`+code+`"}`))
	fmt.Println(res)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		return "", fmt.Errorf("error: status code %d", res.StatusCode)
	}

	// ここでbodyを取得して、それをパースしてアクセストークンを取得する
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	token, ok := result["token"].(string)
	if !ok {
		return "", fmt.Errorf("error: token not found in response")
	}

	return token, nil
}


func GetQiitaArticles(token string) ([]models.QittaList, error) {
	api := "https://qiita.com/api/v2/authenticated_user/items"

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: unexpected status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var items []map[string]interface{}
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var qiitaList []models.QittaList
	for _, item := range items {
		qiita := models.QittaList{}

		// 安全な型変換を行う
		if id, ok := item["id"].(string); ok {
			qiita.Id = id
		}
		if title, ok := item["title"].(string); ok {
			qiita.Title = title
		}
		if body, ok := item["body"].(string); ok {
			qiita.Body = body
		}
		if createdAt, ok := item["created_at"].(string); ok {
			qiita.CreatedAt = createdAt
		}
		if updatedAt, ok := item["updated_at"].(string); ok {
			qiita.UpdatedAt = updatedAt
		}
		if url, ok := item["url"].(string); ok {
			qiita.Url = url
		}
		if private, ok := item["private"].(bool); ok {
			qiita.Private = private
		}
		if likesCount, ok := item["likes_count"].(float64); ok {
			qiita.LikesCount = int(likesCount)
		}

		// タグの処理
		if tags, ok := item["tags"].([]interface{}); ok {
			qiita.Tags = convertToStringSlice(tags)
		}

		qiitaList = append(qiitaList, qiita)
	}

	return qiitaList, nil
}

func convertToStringSlice(interfaces []interface{}) []string {
	strings := make([]string, len(interfaces))
	for i, v := range interfaces {
		if str, ok := v.(map[string]interface{})["name"].(string); ok {
			strings[i] = str
		}
	}
	return strings
}
