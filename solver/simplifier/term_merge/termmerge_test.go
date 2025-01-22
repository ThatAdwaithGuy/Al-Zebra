package termmerge

import (
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestMatchPositive(t *testing.T) {
  testTerm := parser.Addition{
  	Lhs: parser.Variable{Name: "x"},
  	Rhs: parser.Variable{Name: "x"},
  }

  twoTerm := TwoTermMerge{} 

  assert.Equalf(t, parser.Multiplication{
    Lhs: parser.Variable{Name: "x"},
    Rhs: parser.Constant{Value: 2},
  }, twoTerm.GetSimplified(testTerm) , "Should be equal")
}

func TestMatchNegative(t *testing.T) {
  testTerm := parser.Addition{
  	Lhs: parser.Variable{Name: "y"},
  	Rhs: parser.Variable{Name: "x"},
  }

  twoTerm := TwoTermMerge{} 
  assert.Nil(t,twoTerm.GetSimplified(testTerm) , "Should be equal")
}
