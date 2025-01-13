package parser

import (
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
		lhs: Constant{
			Value: 3,
		},
		rhs: Variable{
			Name: "x",
		},
	}

	add := Addition{
		rhs: Constant{Value: 1},
		lhs: mul,
	}

	rhs := Constant{
		Value: 10,
	}

	ast := AST{
		Lhs: add,
		Rhs: rhs,
	}
  ast.Debug()

	assert.Equal(t, &ast, par)
}

