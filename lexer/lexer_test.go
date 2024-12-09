package lexer

import (
	"testing"

	"github.com/nichtsam/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
    let five = 5;
    let ten = 10;
    let add = fn(x, y) {
      x + y;
    };
    let result = add(five, ten);
    !-/*5;
    5 < 10 > 5;
    if (5 < 10) {
      return true;
    } else {
      return false;
    }
    10 == 10;
    10 != 9;
`
	tests := []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "five"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INTERGER, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "ten"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INTERGER, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.L_PAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.R_PAREN, Literal: ")"},
		{Type: token.L_BRACE, Literal: "{"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.R_BRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "result"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.L_PAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "ten"},
		{Type: token.R_PAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.BANG, Literal: "!"},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.SLASH, Literal: "/"},
		{Type: token.ASTERISK, Literal: "*"},
		{Type: token.INTERGER, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INTERGER, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INTERGER, Literal: "10"},
		{Type: token.GT, Literal: ">"},
		{Type: token.INTERGER, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.IF, Literal: "if"},
		{Type: token.L_PAREN, Literal: "("},
		{Type: token.INTERGER, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INTERGER, Literal: "10"},
		{Type: token.R_PAREN, Literal: ")"},
		{Type: token.L_BRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.R_BRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.L_BRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.R_BRACE, Literal: "}"},
		{Type: token.INTERGER, Literal: "10"},
		{Type: token.EQUAL, Literal: "=="},
		{Type: token.INTERGER, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.INTERGER, Literal: "10"},
		{Type: token.NOT_EQUAL, Literal: "!="},
		{Type: token.INTERGER, Literal: "9"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: "\x00"},
	}

	lexer := New(input)

	for i, test := range tests {
		token := lexer.NextToken()
		if token.Type != test.Type {
			t.Fatalf(
				"tests[%d] - token type wrong. Expected=%q, Received=%q",
				i,
				test,
				token,
			)
		}
		if token.Literal != test.Literal {
			t.Fatalf(
				"tests[%d] - token literal wrong. Expected=%q, Received=%q",
				i,
				test,
				token,
			)
		}

	}
}
