package dingadmin

import (
	"github.com/dingdinglz/dingadmin/jwt"
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
		if c.Path() == checkPath(app.config.Prefix, "/login") || c.Path() == checkPath(app.config.Prefix, "/api/login") {
			return c.Next()
		}
		returnUrl := checkPath(app.config.Prefix, "/login")
		if c.Cookies("token", "") == "" {
			return c.Redirect(returnUrl)
		}
		ok, _, _, level := jwt.ParseUserToken(c.Cookies("token", ""), app.tokenKey)
		if !ok {
			return c.Redirect(returnUrl)
		}
		if level < app.config.MinLevel {
			return c.Redirect(returnUrl)
		}
		return c.Next()
	}
}
