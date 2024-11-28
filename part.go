package dingadmin

import "github.com/dingdinglz/dingadmin/tool"

func (app *App) linkParts() {
	if !tool.SringInSlice("user", app.config.DisablePart) {
		app.userDBconnect()
	}
}
