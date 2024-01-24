package utils

import (
	"io"
	"os"
)

type Params struct {
	filename string
	option   []string
	input    string
}

func isOption(arg string) bool {
	return arg[0] == '-'
}

func GetParams() *Params {
	params := new(Params)

	content, err := io.ReadAll(os.Stdin)
	if err == nil {
		params.input = string(content)
	}

	for _, arg := range os.Args[1:] {
		if isOption(arg) {
			params.option = append(params.option, arg)
		} else if content == nil {
			params.filename = arg
		}
	}

	return params
}
