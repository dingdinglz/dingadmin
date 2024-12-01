package dingadmin

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port        int
	Prefix      string
	Name        string
	Title       string
	Author      string
	AuthorLink  string
	Theme       string
	DBtype      string   // DBtype 约定了数据库的类型，例如sqlite或者mysql
	DBsource    string   // DBsource 约定了数据库的连接方法，sqlite下可能为data.db，mysql下可能为user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	DisablePart []string // DisablePart 约定了dingadmin不加载的组件列表，比如user组件
	MinLevel    int      // MinLevel 约定了能够访问admin界面的用户的最小等级
	TokenTime   int      // TokenTime 约定了token能保存的时间长度，单位为分钟
}

var DefaultConfig Config = Config{
	Port:       8080,
	Prefix:     "/admin",
	Name:       "dingadmin",
	Title:      "dingadmin",
	Author:     "dinglz",
	AuthorLink: "https://github.com/dingdinglz",
	Theme:      "./web/admin/",
	DBtype:     "sqlite",
	DBsource:   "data.db",
	MinLevel:   1,
	TokenTime:  60,
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
