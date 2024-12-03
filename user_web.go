package dingadmin

import "github.com/gofiber/fiber/v2"

func (app *App) bindUserRoute() {
	MainRoute := app.Server.Group(app.config.Prefix)
	MainRoute.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", app.GetRenderMap("", c), "layout")
	})
	AdminApiRoute := MainRoute.Group("/api")
	AdminApiRoute.Post("/login", app.ApiLoginRoute())

	g := app.Server.Group(app.config.Prefix, app.AuthMiddleWare())
	g.Get("/user", func(c *fiber.Ctx) error {
		return c.Render("user", app.GetRenderMap("user", c), "layout")
	})
}

func (app *App) bindUserMenus() {
	menuHref := app.config.Prefix
	if menuHref != "/" {
		menuHref += "/"
	}
	menuHref += "user"
	app.AddMenuGroup(MenuGroup{
		Text: "用户",
		Menus: []Menu{
			{
				Name: "user",
				Text: "用户管理",
				Href: menuHref,
			},
		},
	})
}
