package tool

import "strings"

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
