package dingadmin

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
