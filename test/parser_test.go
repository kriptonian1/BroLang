package test

import (
	"testing"

	"github.com/kriptonian1/BroLang/src/ast"
	"github.com/kriptonian1/BroLang/src/lexer"
	"github.com/kriptonian1/BroLang/src/parser"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)

	p := parser.New(l)

	program := p.ParseProgram() // Parses the program from the input string to an AST

	checkParserErrors(t, p) // Checks for parser errors

	if program == nil { // Checks if the program is nil
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 { // Checks if the program has 3 statements
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string // The expected identifier
	}{
		{"x"},      // The first identifier
		{"y"},      // The second identifier
		{"foobar"}, // The third identifier
	}

	for i, tt := range tests {
		stmt := program.Statements[i]                          // Gets the statement
		if !testLetStatement(t, stmt, tt.expectedIdentifier) { // Tests the statement
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram() // Parses the program from the input string to an AST
	checkParserErrors(t, p)     // Checks for parser errors

	if len(program.Statements) != 3 { // Checks if the program has 3 statements
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements { // Loops through the statements
		returnStmt, ok := stmt.(*ast.ReturnStatement) // Type assertion

		if !ok { // Checks if the statement is a return statement
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" { // Checks if the token literal is correct
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram() // Parses the program from the input string to an AST
	checkParserErrors(t, p)     // Checks for parser errors

	if len(program.Statements) != 1 { // Checks if the program has 1 statement
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement) // Type assertion

	if !ok { // Checks if the statement is an expression statement
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier) // Type assertion

	if !ok { // Checks if the expression is an identifier
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" { // Checks if the value of the identifier is correct
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" { // Checks if the token literal of the identifier is correct
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

/*
Tests the let statement
*/
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement) // Type assertion

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name { // Checks if the name of the identifier is correct
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name { // Checks if the token literal of the identifier is correct
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 { // Checks if there are any errors
		return
	}

	t.Errorf("parser has %d errors", len(errors)) // Prints the number of errors

	for _, msg := range errors { // Prints the errors
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
