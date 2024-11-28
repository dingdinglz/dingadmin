package dingadmin

import "time"

type UserModel struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string
	Password  string
	Email     string
	Telephone string
	Level     int
	CreatedAt time.Time
}

func (UserModel) TableName() string {
	return "user"
}

func (app *App) userDBconnect() {
	app.Database.AutoMigrate(&UserModel{})
}
