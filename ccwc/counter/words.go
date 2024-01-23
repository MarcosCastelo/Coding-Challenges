package counter

import "unicode"

func Words(file []byte) (n int) {
	var isWord bool = false
	for _, char := range file {
		if !unicode.IsSpace(rune(char)) {
			if !isWord {
				n++
			}
			isWord = true
		}
		if isWord && unicode.IsSpace(rune(char)) {
			isWord = false
		}
	}
	return
}
