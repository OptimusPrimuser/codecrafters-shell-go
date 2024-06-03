package main

import (
	"os"
	"strconv"
)

var funcMap map[string]func([]string) = map[string]func([]string){
	"exit": exitFunc,
}

func exitFunc(args []string) {
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}
