package counter

import (
	"strings"
	"unicode/utf8"
)

func Chars(file []byte) (n int) {
	n = len(strings.TrimFunc(string(file), func(char rune) bool {
		return !utf8.ValidRune(char)
	}))
	return
}
