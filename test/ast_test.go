package test

import (
	"testing"

	"github.com/kriptonian1/BroLang/src/ast"
	"github.com/kriptonian1/BroLang/src/token"
)

func TestString(t *testing.T) {
	programLet := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
				Value: &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
			},
		},
	}

	programReturn := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
			},
		},
	}

	if programLet.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", programLet.String())
	}

	if programReturn.String() != "return myVar;" {
		t.Errorf("program.String() wrong. got=%q", programReturn.String())
	}
}
