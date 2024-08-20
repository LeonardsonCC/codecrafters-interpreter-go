package main

import (
	"fmt"
	"os"
)

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

func (l *Lox) Match(expected Token) bool {
	if l.current == len(l.source)-1 {
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
