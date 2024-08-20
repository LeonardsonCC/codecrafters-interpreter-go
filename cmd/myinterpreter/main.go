package main

import (
	"fmt"
	"os"
)

type Token string

const (
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

type Lox struct {
	tokens  []Token
	line    int
	current int
	source  []byte
}

func NewLox() *Lox {
	return &Lox{
		tokens: make([]Token, 0, 10),
		line:   1,
	}
}

func (l *Lox) InterpretFile(filename string) []error {
	errs := make([]error, 0, 1)

	f, err := os.ReadFile(filename)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	l.source = f

	for l.current = 0; l.current < len(l.source); l.current++ {
		c := string(l.source[l.current])

		switch c {
		case BREAK_LINE.Token():
			l.line++
		case LEFT_PAREN.Token():
			l.AddToken(LEFT_PAREN)
		case RIGHT_PAREN.Token():
			l.AddToken(RIGHT_PAREN)
		case LEFT_BRACE.Token():
			l.AddToken(LEFT_BRACE)
		case RIGHT_BRACE.Token():
			l.AddToken(RIGHT_BRACE)
		case STAR.Token():
			l.AddToken(STAR)
		case COMMA.Token():
			l.AddToken(COMMA)
		case DOT.Token():
			l.AddToken(DOT)
		case PLUS.Token():
			l.AddToken(PLUS)
		case MINUS.Token():
			l.AddToken(MINUS)
		case SEMICOLON.Token():
			l.AddToken(SEMICOLON)
		case EQUAL.Token():
			if l.Match(EQUAL) {
				l.AddToken(EQUAL_EQUAL)
			} else {
				l.AddToken(EQUAL)
			}
		case BANG.Token():
			if l.Match(EQUAL) {
				l.AddToken(BANG_EQUAL)
			} else {
				l.AddToken(BANG)
			}
		case LESS.Token():
			if l.Match(EQUAL) {
				l.AddToken(LESS_EQUAL)
			} else {
				l.AddToken(LESS)
			}
		case GREATER.Token():
			if l.Match(EQUAL) {
				l.AddToken(GREATER_EQUAL)
			} else {
				l.AddToken(GREATER)
			}
		default:
			err := ErrUnexpectedToken{
				line:  l.line,
				token: c,
			}
			errs = append(errs, err)
			continue
		}

	}
	return errs
}

func (l *Lox) Match(expected Token) bool {
	if l.current == len(l.source) {
		return false
	}

	if string(l.source[l.current+1]) != expected.Token() {
		return false
	}

	l.current++

	return true
}

func (l *Lox) AddToken(r Token) {
	fmt.Printf("%s %s null\n", r.String(), string(r.Token()))
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
