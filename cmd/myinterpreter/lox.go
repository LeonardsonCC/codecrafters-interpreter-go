package main

import (
	"fmt"
	"os"
)

type Token struct {
	tokenType string
	token     string
	literal   *string
}

var (
	EOF Token = Token{
		tokenType: "EOF",
		token:     "",
	}
	BREAK_LINE Token = Token{
		tokenType: "BREAK_LINE",
		token:     "\n",
	}
	LEFT_PAREN Token = Token{
		tokenType: "LEFT_PAREN",
		token:     "(",
	}
	RIGHT_PAREN Token = Token{
		tokenType: "RIGHT_PAREN",
		token:     ")",
	}
	LEFT_BRACE Token = Token{
		tokenType: "LEFT_BRACE",
		token:     "{",
	}
	RIGHT_BRACE Token = Token{
		tokenType: "RIGHT_BRACE",
		token:     "}",
	}
	STAR Token = Token{
		tokenType: "STAR",
		token:     "*",
	}
	DOT Token = Token{
		tokenType: "DOT",
		token:     ".",
	}
	COMMA Token = Token{
		tokenType: "COMMA",
		token:     ",",
	}
	PLUS Token = Token{
		tokenType: "PLUS",
		token:     "+",
	}
	MINUS Token = Token{
		tokenType: "MINUS",
		token:     "-",
	}
	SEMICOLON Token = Token{
		tokenType: "SEMICOLON",
		token:     ";",
	}
	EQUAL Token = Token{
		tokenType: "EQUAL",
		token:     "=",
	}
	BANG Token = Token{
		tokenType: "BANG",
		token:     "!",
	}
	GREATER Token = Token{
		tokenType: "GREATER",
		token:     ">",
	}
	LESS Token = Token{
		tokenType: "LESS",
		token:     "<",
	}
	SLASH Token = Token{
		tokenType: "SLASH",
		token:     "/",
	}
	EQUAL_EQUAL Token = Token{
		tokenType: "EQUAL_EQUAL",
		token:     "==",
	}
	BANG_EQUAL Token = Token{
		tokenType: "BANG_EQUAL",
		token:     "!=",
	}
	LESS_EQUAL Token = Token{
		tokenType: "LESS_EQUAL",
		token:     "<=",
	}
	GREATER_EQUAL Token = Token{
		tokenType: "GREATER_EQUAL",
		token:     ">=",
	}
)

func (t Token) Token() string {
	return t.token
}

func (t Token) String() string {
	return t.tokenType
}

type ErrUnexpectedToken struct {
	line  int
	token string
}

func (e ErrUnexpectedToken) Error() string {
	return fmt.Sprintf("[line %d] Error: Unexpected character: %s\n", e.line, string(e.token))
}

type ErrUnterminedString struct {
	line int
}

func (e ErrUnterminedString) Error() string {
	return fmt.Sprintf("[line %d] Error: Untermined string.\n", e.line)
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
		case `"`:
			var untermined bool
			start := l.current
			end := l.current
			for {
				v := l.Peek()
				if v == "\n" {
					l.line++
					continue
				}
				if l.IsAtEnd() {
					untermined = true
					errs = append(errs, ErrUnterminedString{
						l.line,
					})
					break
				}
				if v == `"` {
					end = l.current + 1
					break
				}
			}

			if untermined {
				break
			}

			literal := string(l.source[start+1 : end-1])
			l.AddToken(Token{
				tokenType: "STRING",
				// TODO: shhhh
				token:   string(l.source[start:end]),
				literal: &literal,
			})
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
	literal := "NULL"
	if r.literal != nil {
		literal = *r.literal
	}

	fmt.Printf("%s %s %s\n", r.String(), string(r.Token()), literal)
	l.tokens = append(l.tokens, r)
}
