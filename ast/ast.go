package ast

import "bytes"

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

func (p *Program) _node() {}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
