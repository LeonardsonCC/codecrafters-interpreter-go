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
	SLASH         Token = "/"
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
	case SLASH:
		return "SLASH"
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
		case " ":
		case "\r":
		case "\t":
			// Ignore whitespace.
			break
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
		case SLASH.Token():
			if l.Match(SLASH) {
				for {
					v := l.Peek()
					if v == "\n" {
						l.line++
						break
					}
					if l.IsAtEnd() {
						break
					}
				}
			} else {
				l.AddToken(SLASH)
			}
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
			if c != "" {
				err := ErrUnexpectedToken{
					line:  l.line,
					token: c,
				}
				errs = append(errs, err)
			}
			continue
		}

	}

	l.AddToken(EOF)

	return errs
}

func (l *Lox) Peek() string {
	if l.IsAtEnd() {
		return ""
	}
	l.current++
	return string(l.source[l.current])
}

func (l *Lox) IsAtEnd() bool {
	return l.current == len(l.source)-1
}

func (l *Lox) Match(expected Token) bool {
	if l.IsAtEnd() {
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
