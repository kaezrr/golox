package lox

import (
	"bufio"
	"fmt"
	"os"
)

func (lox *Lox) RunFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = lox.run(string(bytes))
	if err != nil {
		os.Exit(65)
	}
}

func (lox *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lox.run(line)
	}
}

type Lox struct {
}

func (lox *Lox) run(source string) error {
	// scanner := NewScanner()
	// ...
	return nil
}

func (lox *Lox) reportError(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where, message string) {
	fmt.Println("[line ", line, "] Error", where, ": ", message)
}
