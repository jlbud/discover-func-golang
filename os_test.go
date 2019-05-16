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
