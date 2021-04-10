package main

import (
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
