package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5 返回src的md5编码
func MD5(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
