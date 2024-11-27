package dingadmin

import "testing"

func TestDingAdmin(t *testing.T) {
	cfg := DefaultConfig
	cfg.Theme = "./dingadmin-layui/"
	app := New(cfg)
	app.AddMenu(Menu{Name: "test", Text: "测试菜单", Href: "/test"})
	app.AddMenuGroup(MenuGroup{Text: "测试菜单组", Menus: []Menu{{Name: "test2", Text: "测试菜单2", Href: "/test2"}}})
	app.Serve()
}
