package dingadmin

import (
	"strconv"

	"github.com/dingdinglz/dingadmin/tool"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
	config Config
}

func New(configs ...Config) *App {
	app := &App{}
	app.Server = fiber.New()
	if len(configs) > 0 {
		app.config = configs[0]
	} else {
		app.config = DefaultConfig
	}
	return app
}

func NewWithFiberConfig(serverConfig fiber.Config, configs ...Config) *App {
	app := &App{}
	app.Server = fiber.New(serverConfig)
	if len(configs) > 0 {
		app.config = configs[0]
	} else {
		app.config = DefaultConfig
	}
	return app
}

func (app *App) Serve() error {
	return app.Server.Listen(tool.StringBuilder("0.0.0.0:", strconv.Itoa(app.config.Port)))
}
