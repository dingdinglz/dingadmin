package dingadmin

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port       int
	Prefix     string
	Name       string
	Title      string
	Author     string
	AuthorLink string
	Theme      string
}

var DefaultConfig Config = Config{
	Port:       8080,
	Prefix:     "/admin",
	Name:       "dingadmin",
	Title:      "dingadmin",
	Author:     "dinglz",
	AuthorLink: "https://github.com/dingdinglz",
	Theme:      "./web/admin/",
}

// SaveConfig 保存Config
// p 为保存的文件路径，例如a.json
func (app *App) SaveConfig(p string) {
	c, _ := json.Marshal(&app.config)
	os.WriteFile(p, c, os.ModePerm)
}

// LoadConfig 加载Config
// p 要加载的文件路径
// 返回一个Config
func LoadConfig(p string) (Config, error) {
	b, e := os.ReadFile(p)
	if e != nil {
		return DefaultConfig, e
	}
	var i Config
	e = json.Unmarshal(b, &i)
	if e != nil {
		return DefaultConfig, e
	}
	return i, nil
}
