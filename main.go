package main

import (
	"fmt"
	"golox/lox"
	"os"
)

func main() {
	loxInterpreter := lox.Lox{}
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		loxInterpreter.RunFile(os.Args[1])
	} else {
		loxInterpreter.RunPrompt()
	}
}
