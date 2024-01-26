package main

import (
	"bytes"
)

func GetFilename(args []string) string {
	for _, arg := range args {
		if arg[0] != '-' {
			return arg
		}
	}
	return ""
}

func ByteCounter(input []byte) int {
	return len(input)
}

func LineCounter(input []byte) int {
	return bytes.Count(input, []byte("\n"))
}

func WordCounter(input []byte) int {
	return len(bytes.Fields(input))
}

func CharCounter(input []byte) int {
	return len(bytes.Runes(input))
}
