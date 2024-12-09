package ast

import (
	"bytes"
	"fmt"
)

type Node interface {
	_node()
	String() string
}

type Statement interface {
	_statement()
	Node
}

type Expression interface {
	_expression()
	Node
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Identifier *Identifier
	Expression Expression
}

func (n *LetStatement) _node() {}

func (n *LetStatement) _statement() {}

func (n *LetStatement) String() string {
	return fmt.Sprintf("let %s = %s;", n.Identifier, n.Expression)
}

type ReturnStatement struct {
	Expression Expression
}

func (n *ReturnStatement) _node() {}

func (n *ReturnStatement) _statement() {}

func (n *ReturnStatement) String() string {
	return fmt.Sprintf("return %s;", n.Expression)
}

type Identifier struct {
	Literal string
}

func (n *Identifier) _node() {}

func (n *Identifier) _expression() {}

func (n *Identifier) String() string {
	return n.Literal
}

func (p *Program) _node() {}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
