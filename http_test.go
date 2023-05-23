package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp1(t *testing.T) {
	client := &http.Client{}
	url := "http://mr-api.staging.qingtingfm.com/admins/media/list?page=1&pagesize=10&filter={}"
	request, err := http.NewRequest("GET", url, nil)
	if request == nil {
		t.Fatal("request is nil")
	}

	request.Header.Add("Cookie", "sso_user_id=73696;")

	if err != nil {
		panic(err)
	}
	response, _ := client.Do(request)
	buf, _ := ioutil.ReadAll(response.Body)
	t.Log(string(buf))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Log(err)
		}
	}(response.Body)
}

func TestHttp2(t *testing.T) {
	endPoint := "https://api.openai.com/v1/chat/completions"
	apiKey := "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"

	client := resty.New()
	client.SetProxy("http://127.0.0.1:18081")

	messages := []map[string]string{{"role": "user", "content": "怎么学习英语"}}

	reqBody := make(map[string]interface{})
	reqBody["model"] = "gpt-3.5-turbo"
	reqBody["messages"] = messages
	reqBody["temperature"] = 0.7

	reqJSON, _ := json.Marshal(reqBody)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey)).
		SetBody(reqJSON).
		Post(endPoint)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", resp.String())
}
