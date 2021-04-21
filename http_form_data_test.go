package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

func index(w http.ResponseWriter, r *http.Request) {
	reader, _ := r.MultipartReader()

	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			fmt.Println()
			fmt.Println()
			break
		}
		if p == nil {
			return
		}
		key := p.FormName()
		val, _ := ioutil.ReadAll(p)
		fmt.Printf("key=%s, value=%s", key, string(val))
		fmt.Println()
	}
}

// formdata server
func TestHttpFormataServer(t *testing.T) {
	http.HandleFunc("/china/shanghai", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

// formdata client
func TestHttpFormdataClient(t *testing.T) {
	url := "http://localhost:8080/china/shanghai"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("a_key", "a_val")
	_ = writer.WriteField("b_key", "b_val")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
