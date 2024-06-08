package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type funcStruct struct {
	function func([]string)
	length   int
}

var shellBuiltinFuncMap map[string]funcStruct = map[string]funcStruct{
	"exit": {exitFunc, 1},
	"echo": {echoFunc, -1},
	"type": {typeFunc, 1},
}

var shellBuiltin map[string]bool = map[string]bool{}

func generateShellBuiltIn() {
	for key, _ := range shellBuiltinFuncMap {
		shellBuiltin[key] = true
	}
}

func isExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

var externalShell map[string]string = map[string]string{}

func generateExternalShell() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		entries, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			execPath := path + "/" + entry.Name()
			_, err = exec.LookPath(execPath)
			if err != nil {
				continue
			}
			externalShell[entry.Name()] = execPath
		}
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
	path, ok := externalShell[args[0]]
	if !ok {
		fmt.Printf("%s: not found\n", args[0])
	}
	fmt.Printf("%s is %s\n", args[0], path)
}

func executeExternal(args []string, execPath string) {
	fmt.Println(execPath)
	fmt.Println(args)
	// cmdArray := append([]string{"~/" + execPath}, args...)
	cmd := exec.Command(execPath, args...)
	cmd.Stdout = os.Stdout
	// err := cmd.Run()
	// if err != nil {
	// 	panic(err)
	// }
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(output))
}
