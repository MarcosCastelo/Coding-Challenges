package counter

import "strings"

func Lines(file []byte) (n int) {
	lines := strings.Split(string(file), "\n")
	n = len(lines)
	return
}
