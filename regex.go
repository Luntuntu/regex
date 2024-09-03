package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Token types
const (
	EOF = iota
	IDENTIFIER
	NUMBER
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
	ERR
)

// Token structure
type Token struct {
	Type  int
	Value string
}

// Lexer structure
type Lexer struct {
	input   string
	pos     int
	length  int
	current rune
}

// NewLexer creates a new lexer instance
func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input:  input,
		length: len(input),
	}
	lexer.next() // Initialize the current character
	return lexer
}

// next advances the lexer to the next character
func (l *Lexer) next() {
	if l.pos >= l.length {
		l.current = 0 // End of input
		return
	}
	l.current = rune(l.input[l.pos])
	l.pos++
}

// skipWhitespace skips over any whitespace characters
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.current) {
		l.next()
	}
}

// lexIdentifier parses an identifier
func (l *Lexer) lexIdentifier() Token {
	start := l.pos - 1
	for unicode.IsLetter(l.current) || unicode.IsDigit(l.current) {
		l.next()
	}
	return Token{Type: IDENTIFIER, Value: l.input[start : l.pos-1]}
}

// lexNumber parses a number
func (l *Lexer) lexNumber() Token {
	start := l.pos - 1
	for unicode.IsDigit(l.current) {
		l.next()
	}
	return Token{Type: NUMBER, Value: l.input[start : l.pos-1]}
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	for l.current != 0 {
		switch {
		case unicode.IsSpace(l.current):
			l.skipWhitespace()
		case unicode.IsLetter(l.current):
			return l.lexIdentifier()
		case unicode.IsDigit(l.current):
			return l.lexNumber()
		case l.current == '+':
			l.next()
			return Token{Type: PLUS, Value: "+"}
		case l.current == '-':
			l.next()
			return Token{Type: MINUS, Value: "-"}
		case l.current == '*':
			l.next()
			return Token{Type: MUL, Value: "*"}
		case l.current == '/':
			l.next()
			return Token{Type: DIV, Value: "/"}
		case l.current == '(':
			l.next()
			return Token{Type: LPAREN, Value: "("}
		case l.current == ')':
			l.next()
			return Token{Type: RPAREN, Value: ")"}
		default:
			return Token{Type: ERR, Value: string(l.current)}
		}
	}
	return Token{Type: EOF}
}

// ReadFileContent reads the entire content of a file and returns it as a string
func ReadFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

// main function for demonstration
func main() {
	// Specify the file path
	filePath := "input.txt"

	// Read file content
	content, err := ReadFileContent(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Create a new lexer with the file content
	lexer := NewLexer(content)

	// Process tokens from the lexer
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		fmt.Printf("Token: %+v\n", token)
	}
}
