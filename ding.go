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
	Server         *fiber.App
	config         Config
	dingMenus      []Menu
	dingMenuGroups []MenuGroup
}

func New(configs ...Config) *App {
	app := &App{}
	if len(configs) > 0 {
		app.config = configs[0]
	} else {
		app.config = DefaultConfig
	}
	defaultConfigCheck(&app.config)
	engine := html.New(app.config.Theme, ".html")
	engine.ShouldReload = true
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
	defaultConfigCheck(&app.config)
	engine := html.New(app.config.Theme, ".html")
	serverConfig.Views = engine
	app.Server = fiber.New(serverConfig)
	return app
}

func defaultConfigCheck(cfg *Config) {
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	if cfg.Prefix == "" {
		cfg.Prefix = "/admin"
	}
	if cfg.Name == "" {
		cfg.Name = "dingadmin"
	}
	if cfg.Title == "" {
		cfg.Title = "dingadmin"
	}
	if cfg.Author == "" {
		cfg.Author = "dinglz"
	}
	if cfg.AuthorLink == "" {
		cfg.AuthorLink = "https://github.com/dingdinglz"
	}
	if cfg.Theme == "" {
		cfg.Theme = "./web/admin/"
	}
}

func (app *App) Serve() error {
	bindRouters(app)
	app.Server.Use(logger.New(), recover.New())
	return app.Server.Listen(tool.StringBuilder("0.0.0.0:", strconv.Itoa(app.config.Port)))
}
