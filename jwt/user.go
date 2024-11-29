package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserJWTModel struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Level    int    `json:"level"`
	jwt.RegisteredClaims
}

func MakeUserToken(id uint64, user string, l int, d time.Duration, password string) string {
	var i UserJWTModel
	i.ID = id
	i.Username = user
	i.Level = l
	i.ExpiresAt = jwt.NewNumericDate(time.Now().Add(d))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, i)
	res, _ := token.SignedString([]byte(password))
	return res
}

func ParseUserToken(t string, password string) (bool, uint64, string, int) {
	token, e := jwt.ParseWithClaims(t, &UserJWTModel{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(password), nil
	})
	if e != nil {
		return false, 0, "", 0
	}
	i := token.Claims.(*UserJWTModel)
	return true, i.ID, i.Username, i.Level
}
