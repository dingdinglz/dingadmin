package dingadmin

import "github.com/gofiber/fiber/v2"

func (app *App) GetRenderMap() fiber.Map {
	return fiber.Map{"DingAdminConfig": app.config, "DingAdminMenus": app.dingMenus, "DingAdminMenuGroups": app.dingMenuGroups}
}
