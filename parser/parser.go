package parser

import (
	"github.com/nichtsam/monkey/ast"
	"github.com/nichtsam/monkey/lexer"
	"github.com/nichtsam/monkey/token"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(input string) *Parser {
	lexer := lexer.New(input)
	return &Parser{
		lexer: lexer,

		peekToken: lexer.NextToken(),
	}
}

func (p *Parser) advanceToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	}

	return nil
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.peekToken.Type != token.EOF {
		p.advanceToken()
		if s := p.parseStatement(); s != nil {
			program.Statements = append(program.Statements, s)
		}
	}

	return program
}
