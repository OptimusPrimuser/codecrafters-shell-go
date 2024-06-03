package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.ReplaceAll(cmd, "\n", "")
		fmt.Println(cmd)
		cmdArr := strings.Split(cmd, " ")
		cmdName := cmdArr[0]
		cmdArgs := cmdArr[1:]
		cmdFunc, ok := funcMap[cmdName]
		if !ok {
			fmt.Printf("%s: command not found\n", cmd)
		}
		cmdFunc(cmdArgs)
	}

}
