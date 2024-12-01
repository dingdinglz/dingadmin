package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiMSIsImlzcyI6ImRpbmdhZG1pbiIsImV4cCI6MTczMjg2NDA1N30.92LuCK_oC5g4PLW8ktGbRZhZlWJjCUIHuBVR_wv6nHA

func TestJWT(t *testing.T) {
	type MyC struct {
		Name string `json:"name"`
		jwt.RegisteredClaims
	}
	to, e := jwt.ParseWithClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsImxldmVsIjo1LCJleHAiOjE3MzMwMjY3MzR9.sTbKozBTTSTrnnd5Z2wMyxoEqNLZo4vFPyA2wgSK-_0", &MyC{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("password"), nil
	})
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Println(to.Claims.(*MyC).Name)
	fmt.Println(to.Claims.(*MyC).ExpiresAt.String())
	if !to.Claims.(*MyC).ExpiresAt.Time.After(time.Now()) {
		fmt.Println("已超时")
	}
}

func TestMake(t *testing.T) {
	type MyC struct {
		Name string `json:"name"`
		jwt.RegisteredClaims
	}
	i := MyC{
		"nameTest",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	to := jwt.NewWithClaims(jwt.SigningMethodHS256, i)
	fmt.Println(to.SignedString([]byte("password")))
}

func TestMakeUserToken(t *testing.T) {
	fmt.Println(MakeUserToken(1, "dinglz", 5, time.Minute*2, "test"))
	fmt.Println(ParseUserToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsImxldmVsIjo1LCJleHAiOjE3MzMwMjY3MzR9.sTbKozBTTSTrnnd5Z2wMyxoEqNLZo4vFPyA2wgSK-_0", "test"))
}
