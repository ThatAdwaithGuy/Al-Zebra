package solver

import (
	"testing"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestCarry(t *testing.T) {
	t.SkipNow()
	ex := "3x+1=10"
	lex := lexer.New(ex)
	par, err := parser.Parse(*lex)
	if err != nil {
		t.Fatalf("parsing raised a error %s", err.Error())
	}
	car, err := CarryOperation(par)
	if err != nil {
		t.Fatalf("carry function raised a error %s", err.Error())
	}
	// TODO: 3x and 1 terms are reversed
	lhs := parser.Constant{
		Value: 1,
	}
	var threex parser.Multiplication
	*threex.GetLhs() = parser.Variable{
		Value: "x",
	}
	*threex.GetRhs() = parser.Constant{
		Value: 3,
	}

	var rhs parser.Subtraction
	*threex.GetLhs() = parser.Constant{Value: 10}
	*threex.GetRhs() = threex

	ast := parser.AST{
		Lhs: lhs,
		Rhs: rhs,
	}

	assert.Equal(t, ast, car)
}
