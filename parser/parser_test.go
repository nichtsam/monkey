package parser

import (
	"testing"

	"github.com/nichtsam/monkey/ast"
)

func TestParse(t *testing.T) {
	parser := New("")
	program := parser.Parse()
	checkParseErrors(t, parser)
	checkProgram(t, program, 0)
}

func TestParseLetStatement(t *testing.T) {
	input := `
    let a = b;
    let b = c;
    let c = a;
  `
	parser := New(input)
	program := parser.Parse()
	checkParseErrors(t, parser)
	checkProgram(t, program, 3)

	tests := []*ast.LetStatement{
		{Identifier: &ast.Identifier{Literal: "a"}},
		{Identifier: &ast.Identifier{Literal: "b"}},
		{Identifier: &ast.Identifier{Literal: "c"}},
	}

	for i, test := range tests {
		t.Logf("test[%d]", i)
		s := expectNode[*ast.LetStatement](t, program.Statements[i])
		if s.Identifier.Literal != test.Identifier.Literal {
			t.Fatalf(
				"identifier literal wrong. Expected=%s, Received=%s",
				test.Identifier.Literal,
				s.Identifier.Literal,
			)
		}
	}
}

func TestParseReturnStatement(t *testing.T) {
	input := `
    return a;
    return b;
    return c;
  `
	parser := New(input)
	program := parser.Parse()
	checkParseErrors(t, parser)
	checkProgram(t, program, 3)

	for i := range 3 {
		t.Logf("test[%d]", i)
		expectNode[*ast.ReturnStatement](t, program.Statements[i])
	}
}

func checkProgram(t *testing.T, p *ast.Program, expectedStatementsCount int) {
	if p == nil {
		t.Fatal("Parse() returned nil")
	}

	if len(p.Statements) != expectedStatementsCount {
		t.Fatalf(
			"Program.Statements should contain %d statement(s), received %d instead.",
			expectedStatementsCount,
			len(p.Statements),
		)
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for i, msg := range errors {
		t.Errorf("%d. %q", i, msg)
	}
	t.FailNow()
}

func expectNode[N ast.Node](t *testing.T, node ast.Node) N {
	if expected, ok := node.(N); ok {
		return expected
	}

	t.Fatalf("node type wrong. Expected=%T, Received=%T", *new(N), node)
	return *new(N)
}
