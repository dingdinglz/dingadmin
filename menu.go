package dingadmin

type Menu struct {
	Name string
	Text string
	Href string
}

type MenuGroup struct {
	Menus []Menu
	Text  string
}

func (app *App) AddMenu(m Menu) {
	app.dingMenus = append(app.dingMenus, m)
}

func (app *App) AddMenuGroup(m MenuGroup) {
	app.dingMenuGroups = append(app.dingMenuGroups, m)
}
