package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Token rune

const (
	LEFT_PAREN  Token = '('
	RIGHT_PAREN Token = ')'
	LEFT_BRACE  Token = '{'
	RIGHT_BRACE Token = '}'
	STAR        Token = '*'
	DOT         Token = '.'
	COMMA       Token = ','
	PLUS        Token = '+'
	MINUS       Token = '-'
	SEMICOLON   Token = ';'
	EQUAL       Token = '='
)

var ErrUnexpectedToken = errors.New("unexpected token")

type Lexer struct {
	startExpression bool
	tokens          []Token
	line            int
	// err        error
}

func (l *Lexer) AddToken(r Token) {
	l.tokens = append(l.tokens, r)
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

	// Uncomment this block to pass the first stage

	filename := os.Args[2]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	lexer := new(Lexer)
	lexer.line = 1
	lexer.tokens = make([]Token, 10)

	var exitCode int
	for scanner.Scan() {
		c := rune(scanner.Text()[0])

		switch c {
		case '\n':
			lexer.line++
		case rune(LEFT_PAREN):
			fmt.Println("LEFT_PAREN ( null")
			lexer.AddToken(LEFT_PAREN)
		case rune(RIGHT_PAREN):
			fmt.Println("RIGHT_PAREN ) null")
			lexer.AddToken(RIGHT_PAREN)
		case rune(LEFT_BRACE):
			fmt.Println("LEFT_BRACE { null")
			lexer.AddToken(LEFT_BRACE)
		case rune(RIGHT_BRACE):
			fmt.Println("RIGHT_BRACE } null")
			lexer.AddToken(RIGHT_BRACE)
		case rune(STAR):
			fmt.Println("STAR * null")
			lexer.AddToken(STAR)
		case rune(COMMA):
			fmt.Println("COMMA , null")
			lexer.AddToken(COMMA)
		case rune(DOT):
			fmt.Println("DOT . null")
			lexer.AddToken(DOT)
		case rune(PLUS):
			fmt.Println("PLUS + null")
			lexer.AddToken(PLUS)
		case rune(MINUS):
			fmt.Println("MINUS - null")
			lexer.AddToken(MINUS)
		case rune(SEMICOLON):
			fmt.Println("SEMICOLON ; null")
			lexer.AddToken(SEMICOLON)
		case rune(EQUAL):
			fmt.Println("EQUAL = null")
			lexer.AddToken(EQUAL)
		default:
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", lexer.line, string(c))
			exitCode = 65
			continue
		}

	}
	fmt.Println("EOF  null")
	os.Exit(exitCode)
}
