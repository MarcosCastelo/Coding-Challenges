package main

import (
	"ccwc/counter"
	"ccwc/utils"
	"fmt"
	"os"
)

var (
	filename string
	option   string
	file     []byte
)

func init() {
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if len(os.Args) > 2 {
		option = os.Args[2]
	}
	file = utils.GetFile(filename)
}

func main() {
	switch option {
	case "-c":
		fmt.Println(counter.Bytes(file), filename)
	case "-l":
		fmt.Println(counter.Lines(file), filename)
	case "-w":
		fmt.Println(counter.Words(file), filename)
	case "-m":
		fmt.Println(counter.Chars(file), filename)
	default:
		fmt.Println(counter.Lines(file), counter.Words(file), counter.Bytes(file), filename)
	}

}
