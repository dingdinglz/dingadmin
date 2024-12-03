package dingadmin

import (
	"github.com/gofiber/fiber/v2"
)

func (app *App) GetRenderMap(pageName string, ctx *fiber.Ctx) fiber.Map {
	renderMap := fiber.Map{"DingAdminConfig": app.config, "DingAdminMenus": app.dingMenus, "DingAdminMenuGroups": app.dingMenuGroups, "LayuiThis": pageName}
	username, ok := ctx.Locals("dingadminUser").(string)
	if ok {
		renderMap["DingAdminUser"] = username
	} else {
		renderMap["DingAdminUser"] = ""
	}
	return renderMap
}
