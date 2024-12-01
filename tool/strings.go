package tool

import (
	"math/rand"
	"strings"
	"time"
)

func StringBuilder(s ...string) string {
	var i strings.Builder
	for _, cnt := range s {
		i.WriteString(cnt)
	}
	return i.String()
}

func SringInSlice(s string, ss []string) bool {
	for _, cnt := range ss {
		if cnt == s {
			return true
		}
	}
	return false
}

func GenerateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
