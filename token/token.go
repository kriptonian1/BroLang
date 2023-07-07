package token

type TokenType string // Type of token (e.g. IDENT, INT, ASSIGN, etc.)

// Token struct
type Token struct {
	Type    TokenType // Type of token (e.g. IDENT, INT, ASSIGN, etc.)
	Literal string    // Literal value of token (e.g. add, foobar, 1234567890, etc.)
}

// Token types
const (
	ILLEGAL = "ILLEGAL" // Unknown tokens
	EOF     = "EOF"     // End of File

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234567890

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
