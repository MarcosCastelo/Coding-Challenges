package utils

import (
	"log"
	"os"
)

func GetFile(filename string) (file []byte) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", filename)
	}
	return
}
