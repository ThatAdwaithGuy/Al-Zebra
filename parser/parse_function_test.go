package parser

import (
	"fmt"
	"testing"

	"github.com/al-zebra/lexer"
	"github.com/stretchr/testify/assert"
)

func TestParseSimple(t *testing.T) {
	ex := "3x+1=10"
	lex := lexer.New(ex)
	par, err := Parse(*lex)
	if err != nil {
		t.Error(err.Error())
	}

	mul := Multiplication{
		rhs: Constant{
			Value: 3,
		},
		lhs: Variable{
			Name: "x",
		},
	}

	add := Addition{
		rhs: mul,
		lhs: Constant{
			Value: 1,
		},
	}

	rhs := Constant{
		Value: 10,
	}

	ast := AST{
		Lhs: add,
		Rhs: rhs,
	}

	assert.Equal(t, &ast, par)
}

func TestParseComplex(t *testing.T) {
	ex := "3x/4 = 6"
	lex := lexer.New(ex)
	par, err := Parse(*lex)
	if err != nil {
		t.Error(err.Error())
	}
  fmt.Println("This debug????")
  par.Debug()
}
