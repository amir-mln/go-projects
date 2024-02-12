package parser

import (
	"fmt"

	"github.com/amir-mln/go-projects/pishi/ast"
	"github.com/amir-mln/go-projects/pishi/ast/expressions"
	"github.com/amir-mln/go-projects/pishi/ast/statements"
	"github.com/amir-mln/go-projects/pishi/lexer"
	"github.com/amir-mln/go-projects/pishi/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []error
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) appendPeekError(t token.TokenType) error {
	err := fmt.Errorf("Expected the next token to be %s; Got %s", t, p.peekToken.Type)

	p.errors = append(p.errors, err)

	return err
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Root {
	root := &ast.Root{}
	root.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}

		p.nextToken()
	}

	return root
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *statements.LetStatement {
	stmt := &statements.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &expressions.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	isAsExpected := p.peekTokenIs(t)

	if isAsExpected {
		p.nextToken()
	} else {
		p.appendPeekError(t)
	}

	return isAsExpected
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	p.nextToken()
	p.nextToken()

	return p
}
