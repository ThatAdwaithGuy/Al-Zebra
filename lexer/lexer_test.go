package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexer(t *testing.T) {
	ex := "3x+1=root2(100)"
	lexer := New(ex)
	tokens := lexer.TokenizeAll()
	expected := []Token{
		{
			Type:  NUMBER,
			Value: "3",
		},
		{
			Type:  VARIABLE,
			Value: "x",
		},
		{
			Type:  PLUS,
			Value: "",
		},
		{
			Type:  NUMBER,
			Value: "1",
		},
		{
			Type:  EQUALS,
			Value: "",
		},
		{
			Type:  ROOT,
			Value: "",
		},
		{
			Type:  NUMBER,
			Value: "2",
		},
		{
			Type:  LEFT_PAREN,
			Value: "",
		},
		{
			Type:  NUMBER,
			Value: "100",
		},
		{
			Type:  RIGHT_PAREN,
			Value: "",
		},
		{
			Type:  EOF,
			Value: "",
		},
	}
	assert.Equal(t, expected, tokens, "Incorrect tokenization")
}
