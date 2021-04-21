package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func IsExistFile(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

var VowelMapping = map[string]string{}

func TestRead(t *testing.T) {
	file, err := os.Open("./vowelmapping.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	line := bufio.NewReader(file)
	for {
		line, _, err := line.ReadLine()
		if err == io.EOF || len(line) == 0 {
			break
		}
		strSli := strings.Split(string(line), "##")
		if len(strSli) == 2 {
			VowelMapping[strSli[0]] = strSli[1]
		}
	}

	t.Logf("VowelMapping len=%d", len(VowelMapping))
	t.Log(VowelMapping["<p:uang4>"])
}

// 根据字节数读取file
// 1.16有更好的方式，查看embed_test.go
func TestReadFile(t *testing.T) {
	buf := make([]byte, 10)
	file, _ := os.Open("./embed.txt")
	n, err := io.ReadFull(file, buf)
	if err != nil {
		t.Log(n, err.Error())
	} else {
		t.Log(string(buf))
	}
}
