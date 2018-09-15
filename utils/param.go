package utils

import "strings"

func UpFirstLetter(s string) string {
	if s == "" {
		return ""
	}

	b := []byte(s)
	return strings.ToUpper(string(b[:1])) + string(b[1:])
}
