package lexer

import (
	"testing"

	"github.com/junkboy0315/go-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	expectedResults := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	l := CreateNewLexer(input)

	for i, expectedResult := range expectedResults {
		tok := l.NextToken()

		if tok.Type != expectedResult.Type {
			t.Fatalf("expectedResults[%d] - tokentype wrong. expected=%q, got%q",
				i, expectedResult.Type, tok.Type)
		}

		if tok.Literal != expectedResult.Literal {
			t.Fatalf("expectedResults[%d] - literal wrong. expected=%q, got=%q",
				i, expectedResult.Literal, tok.Literal)
		}
	}
}
