package parser

import (
	"strings"
	"testing"

	"github.com/amir-mln/go-projects/pishi/ast/statements"
	"github.com/amir-mln/go-projects/pishi/lexer"
	"github.com/amir-mln/go-projects/pishi/token"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	for _, msg := range p.Errors() {
		t.Fatalf("parser error: %q", msg)
	}

	tests := []struct{ expectedIdent string }{
		{"x"},
		{"y"},
		{"foobar"},
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf(
			"expected %d Statements from program.Statements; got %d",
			len(tests),
			len(program.Statements),
		)
	}

	for i, c := range tests {
		st := program.Statements[i]

		if st.TokenLiteral() != strings.ToLower(string(token.LET)) {
			t.Errorf("expected the literal value of statement's token to be LET, got %q", st.TokenLiteral())
		}

		letSt, ok := st.(*statements.LetStatement)

		if !ok {
			t.Errorf("expected statement to be of Type ast.LetStatement. got=%T", st)
		}

		if letSt.Name.Value != c.expectedIdent {
			t.Errorf(
				"expected name of let statement identifier to be '%s'. got=%s",
				c.expectedIdent,
				letSt.Name.Value,
			)
		}

		if letSt.Name.TokenLiteral() != c.expectedIdent {
			t.Errorf(
				"letStmt.Name.TokenLiteral() not '%s'. got=%s",
				c.expectedIdent,
				letSt.Name.TokenLiteral(),
			)
		}
	}
}
