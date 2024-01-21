package counter

func checkLetter(char byte) bool {
	return (char >= 65 && char <= 90) || (char >= 97 && char <= 122)
}

func checkDividerChar(char byte) bool {
	return string(char) == " " || string(char) == "\n"
}

func Words(file []byte) (n int) {
	n = 1
	for _, c := range file {
		if string(c) == " " || string(c) == "\n" {
			n++
		}
	}
	return
}
