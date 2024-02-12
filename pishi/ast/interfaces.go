package ast

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	ExpressionNode()
}

type Statement interface {
	Node
	StatementNode()
}
