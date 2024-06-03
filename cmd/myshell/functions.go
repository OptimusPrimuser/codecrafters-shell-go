package main

import (
	"os"
	"strconv"
)

type funcStruct struct {
	function func([]string)
	length   int
}

var funcMap map[string]funcStruct = map[string]funcStruct{
	"exit": {exitFunc, 1},
}

func exitFunc(args []string) {
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}
