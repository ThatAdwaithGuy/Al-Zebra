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
		Lhs: Constant{
			Value: 3,
		},
		Rhs: Variable{
			Value: "x",
		},
	}

	add := Addition{
		Rhs: Constant{Value: 1},
		Lhs: mul,
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


func TestParseSimple2(t *testing.T) {
	ex := "3x+1=root2(10)"
	lex := lexer.New(ex)
	par, err := Parse(*lex)
	if err != nil {
		t.Error(err.Error())
	}

	mul := Multiplication{
		Lhs: Constant{
			Value: 3,
		},
		Rhs: Variable{
			Value: "x",
		},
	}

	add := Addition{
		Rhs: Constant{Value: 1},
		Lhs: mul,
	}

	rhs := Root{
		Lhs: Constant{Value: 2},
		Rhs: Constant{Value: 10},
	}
	ast := AST{
		Lhs: add,
		Rhs: rhs,
	}
  ast.Debug()

	assert.Equal(t, &ast, par)
}
