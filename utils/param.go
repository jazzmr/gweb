package utils

import "strings"

func UpFirstLetter(s string) string {
	b := []byte(s)
	return strings.ToUpper(string(b[:1])) + string(b[1:])
}
