package qitta

import (
	"encoding/json"
	"fmt"
	"io"
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
		return "", fmt.Errorf("Error: token not found in response")
	}

	return token, nil
}