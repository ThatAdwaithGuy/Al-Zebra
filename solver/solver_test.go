package solver

import (
	"testing"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestCarry(t *testing.T) {
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

  threex := parser.NewMultiplication(parser.Variable{
  	Name: "x",
  }, parser.Constant{
  	Value: 3,
  })
  rhs := parser.NewSubtraction(parser.Constant{Value: 10}, threex)
  ast := parser.AST{
  	Lhs: lhs,
  	Rhs: rhs,
  }

  assert.Equal(t, ast, car)
}

