package main

import (
	"bufio"
	"bytes"
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
	file, err := os.Open("./count.txt")
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

func TestStatsCount(t *testing.T) {
	file, _ := os.Open("./count.txt")
	count, _ := StatsCount(file)
	t.Log(count)
}

// statistics file line num
func StatsCount(r io.Reader) (c int, err error) {
	// 32K cache
	buf := make([]byte, 32*1024)
	count := 1
	lineSep := []byte{'\n'}
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
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

//////////////////// Write data to file.
func TestDataWriter(t *testing.T) {
	file, err := os.OpenFile("./tmp/writer.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	writer := bufio.NewWriter(file)
	defer func() {
		_ = writer.Flush()
		_ = file.Close()
	}()

	_, _ = writer.WriteString("a\n")
	_, _ = writer.WriteString("b\n")
}

////////////////////
