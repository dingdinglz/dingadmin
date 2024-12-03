package dingadmin

import (
	"github.com/dingdinglz/dingadmin/jwt"
	"github.com/dingdinglz/dingadmin/tool"
	"github.com/gofiber/fiber/v2"
)

func checkPath(p string, path string) string {
	if p == "/" {
		return path
	}
	return p + path
}

func (app *App) AuthMiddleWare() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if tool.SringInSlice("user", app.config.DisablePart) {
			return c.Next()
		}
		returnUrl := checkPath(app.config.Prefix, "/login")
		if c.Cookies("token", "") == "" {
			return c.Redirect(returnUrl)
		}
		ok, _, username, level := jwt.ParseUserToken(c.Cookies("token", ""), app.tokenKey)
		if !ok {
			return c.Redirect(returnUrl)
		}
		if level < app.config.MinLevel {
			return c.Redirect(returnUrl)
		}
		c.Locals("dingadminUser", username)
		return c.Next()
	}
}
