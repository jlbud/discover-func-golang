package github

import (
	"bufio"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

type DB struct {
	DbType string `ini:"db_type"`
	DbHost string `ini:"db_host"`
	DbPort string `ini:"db_port"`
}

type Redis struct {
	Port string `ini:"port"`
}

func TestParse(t *testing.T) {
	var conf *ini.File
	var err error

	conf, err = ini.Load("test.conf")
	if err != nil {
		fmt.Println(err)
	}
	dbCfg := &DB{}
	err = conf.Section("DB").MapTo(dbCfg)
	if err != nil {
		fmt.Println(err)
	}
	t.Log(dbCfg.DbPort)

	reCfg := &Redis{}
	conf.Section("Redis").MapTo(reCfg)
	t.Log(reCfg.Port)

	str := conf.Section("Redis").Key("port").String()
	t.Log(str)
}

func TestParseDiy(t *testing.T) {
	InitParse("debug")
	nt := Iconfig("Server.level")
	t.Log(nt)
}

var config sync.Map

func InitParse(pattern string) {
	var path string
	pattern = strings.ToLower(pattern)
	if pattern == "debug" || pattern == "test" {
		path = "test.conf"
	} else if pattern == "prod" {
		path = "prod.conf"
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			return
		}
	}()

	reader := bufio.NewReader(f)
	currentSection := ""
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		s := strings.TrimSpace(string(b))
		if s == "" {
			continue
		}

		indexPre := strings.Index(s, "[")
		indexAfter := strings.Index(s, "]")
		if indexPre >= 0 && indexAfter > 0 {
			currentSection = strings.TrimSpace(s[indexPre+1 : indexAfter])
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config.Store(fmt.Sprintf("%s.%s", currentSection, key), value)
	}
}

func Iconfig(key string) string {
	if val, bol := config.Load(key); bol {
		return val.(string)
	} else {
		return ""
	}
}
