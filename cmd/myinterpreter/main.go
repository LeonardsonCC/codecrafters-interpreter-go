package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]

	lox := NewLox()

	var exitCode int
	errs := lox.InterpretFile(filename)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprint(os.Stderr, err.Error())

			if errors.Is(err, ErrUnexpectedToken{}) {
				exitCode = 65
			}
		}
	}
	os.Exit(exitCode)
}
