package tool

import "strings"

func StringBuilder(s ...string) string {
	var i strings.Builder
	for _, cnt := range s {
		i.WriteString(cnt)
	}
	return i.String()
}
