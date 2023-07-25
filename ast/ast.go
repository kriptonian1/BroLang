package ast

import "github.com/kriptonian1/BroLang/token"

// The AST is the abstract syntax tree. It is the data structure that the parser produces.
type Node interface {
	TokenLiteral() string
}

// The Statement and Expression interfaces are used to distinguish between statements and expressions.
// Statements do not produce a value, while expressions do.
type Statement interface {
	Node
	statementNode()
}

// The Expression interface is used to distinguish between expressions and statements.
// Expressions produce a value, while statements do not.
type Expression interface {
	Node
	expressionNode()
}

type Program struct { // The root node of every AST that the parser produces
	Statements []Statement
}

type ReturnStatement struct { // The AST node for the return statement
	Token       token.Token // The token.RETURN token
	ReturnValue Expression  // The expression that is returned
}

func (p *Program) TokenLiteral() string { // Returns the token literal of the first statement in the program
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// The AST node for the let statement
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier // The identifier of the variable
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// The AST node for the identifier
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string      // The value of the identifier
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
