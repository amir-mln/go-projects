package statements

import (
	"github.com/amir-mln/go-projects/pishi/ast"
	"github.com/amir-mln/go-projects/pishi/ast/expressions"
	"github.com/amir-mln/go-projects/pishi/token"
)

type LetStatement struct {
	Token token.Token
	Name  *expressions.Identifier
	Value ast.Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) StatementNode() {}
