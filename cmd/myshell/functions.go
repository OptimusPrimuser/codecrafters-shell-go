package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

var shellBuiltin map[string]bool = map[string]bool{}

func generateShellBuiltIn() {
	for key, _ := range funcMap {
		shellBuiltin[key] = true
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
	_, ok := shellBuiltin[args[0]]
	if ok {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}
	paths := strings.Split(os.Getenv("PATH"), ":")
	// fmt.Println(paths)
	for _, path := range paths {
		execPath := path + "/" + args[0]
		_, err := os.Open(execPath)
		if err != nil {
			continue
			// fmt.Printf("Specified directory %s does not exist\n", path)
		}
		fmt.Printf("%s is %s/%s\n", args[0], path, args[0])
	}

	fmt.Printf("%s: not found\n", args[0])

}
