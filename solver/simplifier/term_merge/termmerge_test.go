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
  assert.Truef(t, twoTerm.Matches(testTerm), "Should match")

  assert.Equalf(t, twoTerm.GetSimplified(testTerm), twoTerm.GetSimplified(parser.Addition{
  	Lhs: parser.Variable{Name: "x"},
  	Rhs: parser.Constant{Value: 2},
  }), "Should be equal")
}

func TestMatchNegative(t *testing.T) {
  testTerm := parser.Addition{
  	Lhs: parser.Variable{Name: "y"},
  	Rhs: parser.Variable{Name: "x"},
  }

  twoTerm := TwoTermMerge{} 
  assert.Falsef(t, twoTerm.Matches(testTerm), "Should not match")
  assert.NotEqualf(t, twoTerm.GetSimplified(testTerm), twoTerm.GetSimplified(parser.Addition{
  	Lhs: parser.Variable{Name: "x"},
  	Rhs: parser.Constant{Value: 2},
  }), "Should not be equal")
}
