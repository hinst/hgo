package hgo

import (
	"strings"
)

const Digits = "0123456789"

func CheckRuneIsDigit(a rune) bool {
	return strings.ContainsRune(Digits, a)
}
