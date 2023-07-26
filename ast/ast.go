package ast

import (
	"bytes"

	"github.com/kriptonian1/BroLang/token"
)

// The AST is the abstract syntax tree. It is the data structure that the parser produces.
type Node interface {
	TokenLiteral() string
	String() string
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

type ExpressionStatement struct { // The AST node for the expression statement
	Token      token.Token // The first token of the expression
	Expression Expression
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
	Value Expression  // The value of the variable
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

/*
The String() method is used to print the AST nodes for debugging purposes.

It is not used to generate the final code.

It returns the string representation of the AST node.
*/
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

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

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (ls *LetStatement) String() string { // Returns the string representation of the let statement
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ") // Writes the token literal of the let statement
	out.WriteString(ls.Name.String())        // Writes the name of the variable
	out.WriteString(" = ")                   // Writes the assignment operator
	if ls.Value != nil {
		out.WriteString(ls.Value.String()) // Writes the value of the variable exa: let x = 5;
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string { // Returns the string representation of the return statement
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (es *ExpressionStatement) String() string { // Returns the string representation of the expression statement
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// this is the string representation of the identifier
func (i *Identifier) String() string { return i.Value }
