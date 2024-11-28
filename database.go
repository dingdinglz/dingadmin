package dingadmin

import (
	"errors"

	"github.com/dingdinglz/dingadmin/tool"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (app *App) connectDatabase() {
	var e error
	switch app.config.DBtype {
	case "sqlite":
		app.Database, e = gorm.Open(sqlite.Open(app.config.DBsource))
	case "mysql":
		app.Database, e = gorm.Open(mysql.Open(app.config.DBsource))
	default:
		e = errors.New("undefined database type")
	}
	if e != nil {
		panic(tool.StringBuilder("dingadmin db connect error!", "\n", e.Error()))
	}
}
