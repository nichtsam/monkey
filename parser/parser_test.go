package parser

import "testing"

func TestParse(t *testing.T) {
	parser := New("")
	program := parser.Parse()
	if program == nil {
		t.Fatal("Parse() returned nil")
	}
}
