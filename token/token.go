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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	// Comparators
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

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
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// Keywords map
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

/*
LookupIdent function takes an identifier and checks if it is a keyword.
If it is a keyword, it returns the keyword's TokenType constant.
If it is not a keyword, it returns the token.IDENT constant.
We use this function when we encounter an identifier in the input.

For example, if we encounter the identifier let, we check if it is a keyword.
If it is, we return the token.LET constant. If it is not, we return the token.IDENT constant.

@params ident string

@return TokenType constant or token.IDENT constant (string)
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
