package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ArgOb struct {
}

var funcMap map[string](func(ArgOb)) = map[string]func(ArgOb){}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.ReplaceAll(cmd, "\n", "")
		_, ok := funcMap[cmd]
		if !ok {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}

}
