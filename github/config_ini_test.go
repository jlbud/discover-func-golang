package github

import (
	"github.com/go-ini/ini"
	"testing"
)

type Config struct {
	Logger  string
	Address string
}

// getConfigFile 读配置文件
func TestGetConfigFile(t *testing.T) {
	iniCfg, err := ini.Load("./config.ini")
	if err != nil {
		t.Error(err)
	}
	sec := iniCfg.Section("config-dev")
	cfg := &Config{
		Logger:  sec.Key("logger").MustString("err.file"),
		Address: sec.Key("address").In("err.file", []string{"", ""}),
	}
	t.Log(cfg.Logger)
	t.Log(cfg.Address)
	t.Logf("%+v", cfg)
}
