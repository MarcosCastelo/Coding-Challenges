package main

import (
	"ccwc/counter"
	"ccwc/utils"
	"fmt"
	"os"
)

func main() {
	option := os.Args[1]
	filename := os.Args[2]

	file := utils.GetFile(filename)

	switch option {
	case "-c":
		fmt.Println(counter.Bytes(file), filename)
	case "-l":
		fmt.Println(counter.Lines(file), filename)
	case "-w":
		fmt.Println(counter.Words(file), filename)
	default:
		fmt.Printf("%v is not a valid option", option)
	}

}
