package dingadmin

import "github.com/dingdinglz/dingadmin/tool"

func (app *App) UserExist(username string) bool {
	var i int64
	app.Database.Model(&UserModel{}).Where("username = ?", username).Count(&i)
	return i != 0
}

// UserLogin 验证用户名和密码
// 返回是否成功，id和level
func (app *App) UserLogin(username string, password string) (bool, uint64, int) {
	if !app.UserExist(username) {
		return false, 0, 0
	}
	var i UserModel
	app.Database.Model(&UserModel{}).Where("username = ?", username).First(&i)
	if tool.MD5(password) == i.Password {
		return true, i.ID, i.Level
	}
	return false, 0, 0
}
