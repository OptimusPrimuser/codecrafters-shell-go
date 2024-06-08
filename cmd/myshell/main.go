package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	generateShellBuiltIn()
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.ReplaceAll(cmd, "\n", "")
		cmdArr := strings.Split(cmd, " ")
		cmdName := cmdArr[0]
		cmdArgs := make([]string, 0)
		if len(cmdArr) > 1 {
			cmdArgs = cmdArr[1:]
		}
		funcOb, ok := funcMap[cmdName]
		if !ok {
			fmt.Printf("%s: command not found\n", cmdName)
			continue
		}
		if funcOb.length != len(cmdArgs) && funcOb.length != -1 {
			fmt.Printf("%s: expected %d got %d args\n", cmdName, funcOb.length, len(cmdArgs))
			continue
		}
		funcOb.function(cmdArgs)

	}

}
