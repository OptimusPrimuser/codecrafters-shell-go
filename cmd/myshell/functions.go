package main

import (
	"fmt"
	"os"
	"strconv"
)

type funcStruct struct {
	function func([]string)
	length   int
}

var funcMap map[string]funcStruct = map[string]funcStruct{
	"exit": {exitFunc, 1},
	"echo": {echoFunc, -1},
	"type": {typeFunc, 1},
}

var typeMap map[string]bool = map[string]bool{}

func generateTypeMap() {
	for key, _ := range funcMap {
		typeMap[key] = true
	}
}

func exitFunc(args []string) {
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}

func echoFunc(args []string) {
	retVal := ""
	for _, arg := range args {
		retVal = retVal + " " + arg
	}
	fmt.Println(retVal[1:])
}

func typeFunc(args []string) {
	_, ok := typeMap[args[0]]
	if ok {
		fmt.Printf("%s is a shell builtin\n", args[0])
	} else {
		fmt.Printf("%s not found\n", args[0])
	}
}
