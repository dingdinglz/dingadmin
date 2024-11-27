package dingadmin

import (
	"strconv"

	"github.com/dingdinglz/dingadmin/tool"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

type App struct {
	Server *fiber.App
	config Config
}

func New(configs ...Config) *App {
	app := &App{}
	if len(configs) > 0 {
		app.config = configs[0]
	} else {
		app.config = DefaultConfig
	}
	engine := html.New(app.config.Theme, ".html")
	app.Server = fiber.New(fiber.Config{Views: engine})
	return app
}

func NewWithFiberConfig(serverConfig fiber.Config, configs ...Config) *App {
	app := &App{}
	if len(configs) > 0 {
		app.config = configs[0]
	} else {
		app.config = DefaultConfig
	}
	engine := html.New(app.config.Theme, ".html")
	serverConfig.Views = engine
	app.Server = fiber.New(serverConfig)
	return app
}

func (app *App) Serve() error {
	app.Server.Use(logger.New(), recover.New())
	return app.Server.Listen(tool.StringBuilder("0.0.0.0:", strconv.Itoa(app.config.Port)))
}
