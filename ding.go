package dingadmin

import (
	"strconv"

	"github.com/dingdinglz/dingadmin/tool"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)

type App struct {
	Server         *fiber.App
	config         Config
	dingMenus      []Menu
	dingMenuGroups []MenuGroup
	Database       *gorm.DB
	tokenKey       string
}

func New(configs ...Config) *App {
	app := &App{}
	if len(configs) > 0 {
		app.config = configs[0]
	} else {
		app.config = DefaultConfig
	}
	defaultConfigCheck(&app.config)
	app.connectDatabase()
	engine := html.New(app.config.Theme, ".html")
	engine.ShouldReload = true
	app.Server = fiber.New(fiber.Config{Views: engine})
	app.linkParts()
	app.tokenKey = tool.GenerateRandomString(20)
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
	app.connectDatabase()
	engine := html.New(app.config.Theme, ".html")
	serverConfig.Views = engine
	app.Server = fiber.New(serverConfig)
	app.linkParts()
	app.tokenKey = tool.GenerateRandomString(20)
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
	if cfg.DBtype == "" {
		cfg.DBtype = "sqlite"
	}
	if cfg.DBsource == "" {
		cfg.DBsource = "data.db"
	}
	if cfg.MinLevel == 0 {
		cfg.MinLevel = 1
	}
	if cfg.TokenTime == 0 {
		cfg.TokenTime = 60
	}
}

func (app *App) Serve() error {
	bindRouters(app)
	app.Server.Use(logger.New(), recover.New())
	return app.Server.Listen(tool.StringBuilder("0.0.0.0:", strconv.Itoa(app.config.Port)))
}
