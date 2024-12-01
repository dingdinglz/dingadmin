package dingadmin

import "testing"

func TestDingAdmin(t *testing.T) {
	c, _ := LoadConfig("test.json")
	app := New(c)
	app.NewUser("admin", "admin", "test@qq.com", "15156115515", 5)
	app.AddMenu(Menu{Name: "test", Text: "测试菜单", Href: "/test"})
	app.AddMenu(Menu{Name: "about", Text: "关于", Href: "https://github.com/dingdinglz"})
	app.AddMenuGroup(MenuGroup{Text: "测试菜单组", Menus: []Menu{{Name: "test2", Text: "测试菜单2", Href: "/test2"}}})
	app.Serve()
}
