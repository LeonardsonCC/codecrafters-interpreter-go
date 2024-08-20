package main

import (
	"fmt"
	"os"
)

type Token string

const (
	EOF           Token = ""
	BREAK_LINE    Token = "\n"
	LEFT_PAREN    Token = "("
	RIGHT_PAREN   Token = ")"
	LEFT_BRACE    Token = "{"
	RIGHT_BRACE   Token = "}"
	STAR          Token = "*"
	DOT           Token = "."
	COMMA         Token = ","
	PLUS          Token = "+"
	MINUS         Token = "-"
	SEMICOLON     Token = ";"
	EQUAL         Token = "="
	BANG          Token = "!"
	GREATER       Token = ">"
	LESS          Token = "<"
	EQUAL_EQUAL   Token = "=="
	BANG_EQUAL    Token = "!="
	LESS_EQUAL    Token = "<="
	GREATER_EQUAL Token = ">="
)

func (t Token) Token() string {
	return string(t)
}

func (t Token) String() string {
	switch t {
	case EOF:
		return "EOF"
	case BREAK_LINE:
		return "BREAK_LINE"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case STAR:
		return "STAR"
	case DOT:
		return "DOT"
	case COMMA:
		return "COMMA"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SEMICOLON:
		return "SEMICOLON"
	case EQUAL:
		return "EQUAL"
	case BANG:
		return "BANG"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	default:
		return ""
	}
}

type ErrUnexpectedToken struct {
	line  int
	token string
}

func (e ErrUnexpectedToken) Error() string {
	return fmt.Sprintf("[line %d] Error: Unexpected character: %s\n", e.line, string(e.token))
}

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

	errs := lox.InterpretFile(filename)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprint(os.Stderr, err.Error())
		}
		os.Exit(65)
	}
}
