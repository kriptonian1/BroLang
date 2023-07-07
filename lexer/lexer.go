package lexer

import "github.com/kriptonian1/BroLang/token"

type Lexer struct {
	input        string // Input to be tokenized
	position     int    // Current position in input (points to current char)
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // Current char under examination
}

/*
New creates a new lexer instance

@param input string - Input to be tokenized

@return *Lexer - A new lexer instance
*/
func New(input string) *Lexer {
	l := &Lexer{input: input} // Create a new lexer instance with the input
	l.readChar()              // Read the first character in the input
	return l
}

/*
readChar reads the next character in the input and advances the position in the input string
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition] // Get the current character
	}
	l.position = l.readPosition // Update the position to the read position (current position)
	l.readPosition += 1         // Increment the read position
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = "" // End of file
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() // Read the identifier
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

/*
newToken creates a new token instance

@param tokenType token.TokenType - Type of token (e.g. IDENT, INT, ASSIGN, etc.)

@param ch byte - Character to be tokenized

@return token.Token - A new token instance
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier reads an identifier and advances the position until it encounters a non-letter character
func (l *Lexer) readIdentifier() string {
	position := l.position // Save the current position
	for isLetter(l.ch) {
		l.readChar() // Read the next character
	}
	return l.input[position:l.position] // Return the identifier
}

/*
isLetter checks if a character is a letter

@param ch byte - Character to be checked

@return bool - True if the character is a letter, false otherwise
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
