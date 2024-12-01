package tool

import (
	"fmt"
	"testing"
)

func TestGenerateString(t *testing.T) {
	fmt.Println(GenerateRandomString(50))
}
