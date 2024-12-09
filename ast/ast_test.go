package ast

import "testing"

func TestString(t *testing.T) {
	ast := Program{[]Statement{
		&LetStatement{
			Identifier: &Identifier{Literal: "a"},
			Expression: &Identifier{Literal: "b"},
		},
		&LetStatement{
			Identifier: &Identifier{Literal: "b"},
			Expression: &Identifier{Literal: "c"},
		},
		&LetStatement{
			Identifier: &Identifier{Literal: "c"},
			Expression: &Identifier{Literal: "a"},
		},
	}}

	expected := "let a = b;let b = c;let c = a;"
	received := ast.String()

	if received != expected {
		t.Fatalf("String() wrong.\nExpected=%q\nReceived=%q", expected, received)
	}
}
