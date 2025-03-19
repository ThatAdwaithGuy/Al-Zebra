package simplifier

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

	twoTerm, _ := TwoTermMerge(testTerm)

	assert.Equalf(t, parser.Multiplication{
		Lhs: parser.Variable{Name: "x"},
		Rhs: parser.Constant{Value: 2},
	}, twoTerm, "Should be equal")
}

func TestMatchNegative(t *testing.T) {
	testTerm := parser.Addition{
		Lhs: parser.Variable{Name: "y"},
		Rhs: parser.Variable{Name: "x"},
	}

	twoTerm, _ := TwoTermMerge(testTerm)
	assert.Nil(t, twoTerm, "Should be equal")
}
