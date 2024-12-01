package dingadmin

import "github.com/gofiber/fiber/v2"

func bindRouters(app *App) {
	MainRoute := app.Server.Group(app.config.Prefix, app.AuthMiddleWare())
	MainRoute.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", app.GetRenderMap(), "layout")
	})
	MainRoute.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", app.GetRenderMap(), "layout")
	})

	AdminApiRoute := MainRoute.Group("/api")
	AdminApiRoute.Post("/login", app.ApiLoginRoute())
}
