package lexer

import (
	"fmt"
	"unicode"
)

// TokenType represents the type of token
type TokenType int

const (
	NUMBER   TokenType = iota
	PLUS               // 1
	MINUS              // 2
	MULTIPLY           // 3
	DIVIDE             // 4
	ROOT               // 5
  EXPONENTIATION
	EQUALS
	LEFT_PAREN
	RIGHT_PAREN
	VARIABLE 
	EOF
)

// String returns the string representation of the token type
func (t TokenType) String() string {
	return [...]string{
		"NUMBER",
		"PLUS",
		"MINUS",
		"MULTIPLY",
		"DIVIDE",
		"ROOT",
    "EXPONENTIATION",
		"EQUALS",
		"LEFT_PAREN",
		"RIGHT_PAREN",
		"VARIABLE", 
		"EOF",
	}[t]
}

// Token represents a lexical token
type Token struct {
	Type  TokenType
	Value string
}

// String returns a string representation of the token
func (t Token) String() string {
	if t.Value != "" {
		return fmt.Sprintf("Token(%v, %s)", t.Type, t.Value)
	}
	return fmt.Sprintf("Token(%v)", t.Type)
}

// Lexer performs lexical analysis
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// peekChar returns the next character without advancing
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// skipWhitespace skips any whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch != 0 && unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

// readNumber reads a number (integer or decimal)
func (l *Lexer) readNumber() string {
	position := l.position
	for l.ch != 0 && (unicode.IsDigit(rune(l.ch)) || l.ch == '.') {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readVariable reads a single alphabetic character for variables
func (l *Lexer) readVariable() string {
	position := l.position
	l.readChar() // Read the alphabetic character
	return l.input[position:l.position]
}

// NextToken returns the next token in the input
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = Token{Type: PLUS}
	case '-':
		tok = Token{Type: MINUS}
	case '*':
		tok = Token{Type: MULTIPLY}
	case '/':
		tok = Token{Type: DIVIDE}
	case '^':
		tok = Token{Type: EXPONENTIATION}
	case '=':
		tok = Token{Type: EQUALS}
	case '(':
		tok = Token{Type: LEFT_PAREN}
	case ')':
		tok = Token{Type: RIGHT_PAREN}
	case 0:
		tok = Token{Type: EOF}
	default:
		if unicode.IsDigit(rune(l.ch)) {
			numStr := l.readNumber()
			return Token{Type: NUMBER, Value: numStr}
		} else if l.ch == 'r' && len(l.input[l.position:]) >= 4 {
			// Check for 'root' keyword
			if l.input[l.position:l.position+4] == "root" {
				l.readPosition = l.position + 4
				l.readChar()
				return Token{Type: ROOT}
			}
		} else if unicode.IsLetter(rune(l.ch)) { // Check if it's a variable (alphabetic character)
			varStr := l.readVariable()
			return Token{Type: VARIABLE, Value: varStr}
		}
		tok = Token{Type: EOF, Value: fmt.Sprintf("illegal character: %c", l.ch)}
	}

	l.readChar()
	return tok
}

// TokenizeAll returns all tokens in the input
func (l *Lexer) TokenizeAll() []Token {
	var tokens []Token
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
	}
	return tokens
}
