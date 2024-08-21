package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Token struct {
	tokenType string
	token     string
	literal   any
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
	AND Token = Token{
		tokenType: "AND",
		token:     "and",
	}
	OR Token = Token{
		tokenType: "OR",
		token:     "or",
	}
	CLASS Token = Token{
		tokenType: "CLASS",
		token:     "class",
	}
	ELSE Token = Token{
		tokenType: "ELSE",
		token:     "else",
	}
	FALSE Token = Token{
		tokenType: "FALSE",
		token:     "false",
	}
	FOR Token = Token{
		tokenType: "FOR",
		token:     "for",
	}
	FUN Token = Token{
		tokenType: "FUN",
		token:     "fun",
	}
	IF Token = Token{
		tokenType: "IF",
		token:     "if",
	}
	NIL Token = Token{
		tokenType: "NIL",
		token:     "nil",
	}
	PRINT Token = Token{
		tokenType: "PRINT",
		token:     "print",
	}
	RETURN Token = Token{
		tokenType: "RETURN",
		token:     "return",
	}
	SUPER Token = Token{
		tokenType: "SUPER",
		token:     "super",
	}
	THIS Token = Token{
		tokenType: "THIS",
		token:     "this",
	}
	TRUE Token = Token{
		tokenType: "TRUE",
		token:     "true",
	}
	VAR Token = Token{
		tokenType: "VAR",
		token:     "var",
	}
	WHILE Token = Token{
		tokenType: "WHILE",
		token:     "while",
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

type ErrUnterminatedString struct {
	line int
}

func (e ErrUnterminatedString) Error() string {
	return fmt.Sprintf("[line %d] Error: Unterminated string.\n", e.line)
}

type Lox struct {
	tokens   []Token
	line     int
	current  int
	source   []byte
	keywords map[string]Token
}

func NewLox() *Lox {
	kw := map[string]Token{
		AND.Token():    AND,
		OR.Token():     OR,
		CLASS.Token():  CLASS,
		ELSE.Token():   ELSE,
		FALSE.Token():  FALSE,
		FOR.Token():    FOR,
		FUN.Token():    FUN,
		IF.Token():     IF,
		NIL.Token():    NIL,
		PRINT.Token():  PRINT,
		RETURN.Token(): RETURN,
		SUPER.Token():  SUPER,
		THIS.Token():   THIS,
		TRUE.Token():   TRUE,
		VAR.Token():    VAR,
		WHILE.Token():  WHILE,
	}

	return &Lox{
		tokens:   make([]Token, 0, 10),
		line:     1,
		keywords: kw,
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

	for l.current = 0; l.current < len(l.source); l.Advance() {
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
					v, ok := l.Peek()
					if !ok {
						break
					}
					l.Advance()

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
				v, ok := l.Peek()
				if !ok || l.IsAtEnd() {
					untermined = true
					errs = append(errs, ErrUnterminatedString{
						l.line,
					})
					break
				}

				l.Advance()
				if v == "\n" {
					l.line++
					continue
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
				token:     string(l.source[start:end]),
				literal:   &literal,
			})
		default:
			if unicode.IsDigit(rune(c[0])) {
				l.number()
			} else if l.isAlpha(c) {
				l.identifier()
			} else {
				err := ErrUnexpectedToken{
					line:  l.line,
					token: c,
				}
				errs = append(errs, err)
				continue
			}
		}

	}

	l.AddToken(EOF)

	return errs
}

func (l *Lox) Advance() {
	l.current++
}

func (l *Lox) Peek() (string, bool) {
	if l.IsAtEnd() {
		return "", false
	}
	return string(l.source[l.current+1]), true
}

func (l *Lox) PeekNext() (string, bool) {
	if l.current+2 >= len(l.source) {
		return "", false
	}
	return string(l.source[l.current+2]), true
}

func (l *Lox) IsAtEnd() bool {
	return l.current >= len(l.source)-1
}

func (l *Lox) Match(expected Token) bool {
	if l.IsAtEnd() {
		return false
	}

	if string(l.source[l.current+1]) != expected.Token() {
		return false
	}

	l.Advance()

	return true
}

func (l *Lox) AddToken(r Token) {
	literal := "null"
	if r.literal != nil {
		switch v := r.literal.(type) {
		case *string:
			literal = fmt.Sprintf("%v", *v)
		case float64:
			if v == float64(int(v)) {
				literal = fmt.Sprintf("%.1f", v)
			} else {
				literal = fmt.Sprintf("%g", v)
			}
		default:
			literal = fmt.Sprintf("%v", v)
		}
	}

	fmt.Printf("%s %s %s\n", r.String(), string(r.Token()), literal)
	l.tokens = append(l.tokens, r)
}

func (l *Lox) number() {
	start := l.current
	end := l.current
	for {
		v, ok := l.Peek()
		if !ok {
			break
		}

		if !unicode.IsDigit(rune(v[0])) {
			break
		}

		l.Advance()
	}

	if v2, ok := l.Peek(); ok && v2 == DOT.Token() {
		if v2, ok := l.PeekNext(); ok {
			if unicode.IsDigit(rune(v2[0])) {
				l.Advance()
			}
			for {
				v, ok := l.Peek()
				if !ok {
					break
				}

				if !unicode.IsDigit(rune(v[0])) {
					break
				}
				l.Advance()
			}
		}
	}

	end = l.current + 1
	text := string(l.source[start:end])
	literal, err := strconv.ParseFloat(text, 64)
	if err != nil {
		// TODO: emit error
	}

	l.AddToken(Token{
		tokenType: "NUMBER",
		token:     text,
		literal:   literal,
	})
}

func (l *Lox) isAlpha(c string) bool {
	return unicode.IsLetter(rune(c[0])) || c == "_"
}

func (l *Lox) identifier() {
	start := l.current
	end := l.current

	for {
		v, ok := l.Peek()
		if !ok {
			end = l.current + 1
			break
		}

		if l.isAlpha(v) {
			l.Advance()
			continue
		}

		end = l.current + 1
		break
	}

	term := string(l.source[start:end])

	tokenType := "IDENTIFIER"
	if v, ok := l.keywords[term]; ok {
		tokenType = v.String()
	}

	l.AddToken(Token{
		tokenType: tokenType,
		token:     term,
		// literal:   &literal,
	})
}
