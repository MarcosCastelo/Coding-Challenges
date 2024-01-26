package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	output   string
	fileData []byte
)

func main() {
	byteFlag := flag.Bool("c", false, "output number of bytes")
	lineFlag := flag.Bool("l", false, "output number of lines")
	wordFlag := flag.Bool("w", false, "output number of lines")
	charFlag := flag.Bool("m", false, "output number of lines")
	flag.Parse()

	filename := GetFilename(flag.Args())
	if len(filename) > 0 {
		file, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", filename)
			return
		}
		fileData = file
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fileData = stdin
	}

	switch {
	case *byteFlag:
		output += fmt.Sprintf("%d ", ByteCounter(fileData))
	case *lineFlag:
		output += fmt.Sprintf("%d ", LineCounter(fileData))
	case *wordFlag:
		output += fmt.Sprintf("%d ", WordCounter(fileData))
	case *charFlag:
		output += fmt.Sprintf("%d ", CharCounter(fileData))
	default:
		output += fmt.Sprintf("%d %d %d ", ByteCounter(fileData), LineCounter(fileData), WordCounter(fileData))
	}

	if len(filename) > 0 {
		output += filename
	}

	fmt.Println(output)
}
