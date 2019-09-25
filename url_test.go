package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	//a := url.QueryEscape("https://blog.csdn.net/u010214802")
	a := url.PathEscape("https://blog.csdn.net/u010214802")
	fmt.Println(a)
}
