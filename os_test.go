package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOsname(t *testing.T) {
	name, _ := os.Hostname()
	fmt.Println(name)
}

func TestOsPWD(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
}
