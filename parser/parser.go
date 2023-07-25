package parser

import (
	"errors"
	"fmt"

	"github.com/kriptonian1/BroLang/ast"
	"github.com/kriptonian1/BroLang/lexer"
	"github.com/kriptonian1/BroLang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // Current token
	peekToken token.Token // Next token
	errors    []error     // Errors
}

func New(l *lexer.Lexer) *Parser { // Creates a new parser
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	// Read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []error { // Returns the errors
	return p.errors
}

func (p *Parser) nextToken() { // Advances the tokens
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program { // Parses the program
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF { // Loops until the end of the file
		stmt := p.parseStatement() // Parses the statement
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()

	}
	return program
}

func (p *Parser) parseStatement() ast.Statement { // Parses the statement
	switch p.curToken.Type {
	case token.LET: // If the token is a let token
		return p.parseLetStatement() // Parses the let statement
	case token.RETURN: // If the token is a return token
		return p.parseReturnStatement() // Parses the return statement
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken} // Creates a new let statement

	if !p.expectPeek(token.IDENT) { // Checks if the next token is an identifier
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // Sets the name of the identifier

	if !p.expectPeek(token.ASSIGN) { // Checks if the next token is an assign token
		return nil
	}

	//TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) { // Loops until the token is a semicolon
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken} // Creates a new return statement

	p.nextToken()

	//TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) { // Loops until the token is a semicolon
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool { // Checks if the current token is a certain token
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool { // Checks if the next token is a certain token
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	err := errors.New(msg)
	p.errors = append(p.errors, err)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
