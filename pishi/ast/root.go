package ast

type Root struct {
	Statements []Statement
}

func (p *Root) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
