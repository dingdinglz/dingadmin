package dingadmin

func (app *App) UserExist(username string) bool {
	var i int64
	app.Database.Model(&UserModel{}).Where("username = ?", username).Count(&i)
	return i != 0
}
