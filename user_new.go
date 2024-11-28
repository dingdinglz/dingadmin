package dingadmin

import (
	"errors"

	"github.com/dingdinglz/dingadmin/tool"
)

func (app *App) NewUser(Username string, Password string, Email string, Telephone string, Level int) error {
	if app.UserExist(Username) {
		return errors.New("user existed")
	}
	var i UserModel
	i.Username = Username
	i.Password = tool.MD5(Password)
	i.Email = Email
	i.Telephone = Telephone
	i.Level = Level
	app.Database.Create(&i)
	return nil
}
