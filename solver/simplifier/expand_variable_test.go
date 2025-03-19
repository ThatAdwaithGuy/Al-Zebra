package simplifier

import (
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestExpandVariableConstantPositive(t *testing.T) {
	// 3(x+2)
	actual := parser.Multiplication{
		Lhs: parser.Constant{Value: 3},
		Rhs: parser.Addition{
			Lhs: parser.Variable{Name: "x"},
			Rhs: parser.Constant{Value: 2},
		},
	}

	expected := parser.Addition{
		Lhs: parser.Multiplication{
			Lhs: parser.Constant{Value: 3},
			Rhs: parser.Variable{Name: "x"},
		},
		Rhs: parser.Multiplication{
			Lhs: parser.Constant{Value: 3},
			Rhs: parser.Constant{Value: 2},
		},
	}

	pattern, _ := ExpandVariableConstantOperation(actual)
	assert.Equal(t, expected, pattern)
}

func TestExpandVariableConstantNegative(t *testing.T) {
	// 3(x+2)
	actual := parser.Addition{
		Lhs: parser.Constant{Value: 3},
		Rhs: parser.Addition{
			Lhs: parser.Variable{Name: "x"},
			Rhs: parser.Constant{Value: 2},
		},
	}

	pattern, _ := ExpandVariableConstantOperation(actual)
	assert.Nil(t, pattern)
}
