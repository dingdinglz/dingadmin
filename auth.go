package dingadmin

import (
	"time"

	"github.com/dingdinglz/dingadmin/jwt"
	"github.com/gofiber/fiber/v2"
)

func (app *App) ApiLoginRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.FormValue("username", "") == "" || c.FormValue("password", "") == "" {
			return c.JSON(fiber.Map{"code": -1, "message": "参数不全！"})
		}
		ok, id, level := app.UserLogin(c.FormValue("username"), c.FormValue("password"))
		if !ok {
			return c.JSON(fiber.Map{"code": 1, "message": "密码错误或用户不存在！"})
		}
		token := jwt.MakeUserToken(id, c.FormValue("username"), level, time.Minute*time.Duration(app.config.TokenTime), app.tokenKey)
		c.Cookie(&fiber.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(time.Minute * time.Duration(app.config.TokenTime)),
		})
		return c.JSON(fiber.Map{"code": 0, "message": "登录成功！", "redirect": app.config.Prefix})
	}
}
