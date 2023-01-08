package fuzzing

import (
	"errors"
	"unicode/utf8"
)

func Reverse1(str string) string {
	bytes := []byte(str)
	for i, j := 0, len(bytes)-1; i < len(bytes)/2; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)
}

func Reverse2(str string) (string, error) {
	if !utf8.ValidString(str) {
		return str, errors.New("input is not valid UTF-8")
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}
