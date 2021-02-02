package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// https请求去掉证书验证
func TestHttpServer(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// HTTP中间件
func TestHttpMiddleware(t *testing.T) {
	http.Handle("/hello", timeMiddleWare(http.HandlerFunc(middleWareHello)))
	http.Handle("/world", timeMiddleWare(http.HandlerFunc(middleWareWorld)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func middleWareHello(wr http.ResponseWriter, r *http.Request) {
	_, _ = wr.Write([]byte("hello"))
}

func middleWareWorld(wr http.ResponseWriter, r *http.Request) {
	_, _ = wr.Write([]byte("world"))
}

// 中间件通过包装handler再返回一个新的handler
func timeMiddleWare(next http.Handler) http.Handler {
	// TODO 没明白
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		// 中间件做了打印时间的事情
		timS := time.Now()
		next.ServeHTTP(wr, r)
		timeE := time.Since(timS)
		fmt.Println(timeE)
	})
}
