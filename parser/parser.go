package parser

import (
	"fmt"

	"github.com/nichtsam/monkey/ast"
	"github.com/nichtsam/monkey/lexer"
	"github.com/nichtsam/monkey/token"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(input string) *Parser {
	lexer := lexer.New(input)
	return &Parser{
		lexer: lexer,

		peekToken: lexer.NextToken(),
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) error(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) advanceToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expectNextToken(expected token.TokenType) bool {
	if p.peekToken.Type != expected {
		p.error(fmt.Sprintf("expected token type %s, got %s instead", expected, p.peekToken.Type))
		return false
	}

	p.advanceToken()
	return true
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		if s := p.parseLetStatement(); s != nil {
			return s
		}
	}

	return nil
}

func (p *Parser) parseExpression() ast.Expression {
	// TODO: parse Expression
	for p.curToken.Type != token.SEMICOLON {
		p.advanceToken()
	}

	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	s := &ast.LetStatement{}

	if ok := p.expectNextToken(token.IDENTIFIER); !ok {
		return nil
	}
	s.Identifier = &ast.Identifier{Literal: p.curToken.Literal}

	if ok := p.expectNextToken(token.ASSIGN); !ok {
		return nil
	}

	s.Expression = p.parseExpression()

	return s
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
