package simplifier

import (
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestHasVariable(t *testing.T) {
	lhs := parser.Addition{
		Lhs: parser.Multiplication{
			Lhs: parser.Constant{
				Value: 3,
			},
			Rhs: parser.Variable{
				Value: "x",
			},
		},
		Rhs: parser.Constant{
			Value: 1,
		},
	}

	// rhs := parser.Constant{
	// 	Value: 10,
	// }
	//
	// ast := parser.AST{
	// 	Lhs: lhs,
	// 	Rhs: rhs,
	// }

	assert.Equal(t, true, hasVariable(lhs))
}

func TestBasicAlzebraPass(t *testing.T) {
	lhs := parser.Addition{
		Lhs: parser.Multiplication{
			Lhs: parser.Constant{
				Value: 3,
			},
			Rhs: parser.Variable{
				Value: "x",
			},
		},
		Rhs: parser.Constant{
			Value: 1,
		},
	}

	rhs := parser.Constant{
		Value: 10,
	}

	ast := parser.AST{
		Lhs: lhs,
		Rhs: rhs,
	}

	pattern := BasicAlgebra{}

	assert.Equal(t, true, pattern.IsValid(&ast), "Validation is incorrect. expected true but got false")
}

func TestBasicAlzebraFailExp(t *testing.T) {
	lhs := parser.Addition{
		Lhs: parser.Exponentiation{
			Lhs: parser.Constant{
				Value: 3,
			},
			Rhs: parser.Variable{
				Value: "x",
			},
		},
		Rhs: parser.Constant{
			Value: 1,
		},
	}

	rhs := parser.Constant{
		Value: 10,
	}

	ast := parser.AST{
		Lhs: lhs,
		Rhs: rhs,
	}

	pattern := BasicAlgebra{}

	assert.Equal(t, false, pattern.IsValid(&ast), "Validation is incorrect. expected false but got true")
}

func TestBasicAlzebraFailRoot(t *testing.T) {
	lhs := parser.Addition{
		Lhs: parser.Root{
			Lhs: parser.Subtraction{
				Lhs: parser.Variable{
					Value: "x",
				},
				Rhs: parser.Constant{
					Value: 4,
				},
			},
			Rhs: parser.Variable{
				Value: "x",
			},
		},
		Rhs: parser.Constant{
			Value: 1,
		},
	}

	rhs := parser.Constant{
		Value: 10,
	}

	ast := parser.AST{
		Lhs: lhs,
		Rhs: rhs,
	}

	pattern := BasicAlgebra{}

	assert.Equal(t, false, pattern.IsValid(&ast), "Validation is incorrect. expected false but got true")
}

func TestTransferFunctionVariableConstantNil(t *testing.T) {
	astContant := parser.AST{
		Lhs: parser.Constant{Value: 1},
		Rhs: parser.Constant{Value: 3},
	}
	assert.Nil(t, transferTermRtoL(&astContant, false))

	astVariable := parser.AST{
		Lhs: parser.Variable{Value: "x"},
		Rhs: parser.Variable{Value: "y"},
	}

	assert.Nil(t, transferTermRtoL(&astVariable, false))
}

func TestTransferFunctionVariableConstantPass(t *testing.T) {
	ast := parser.AST{
		Lhs: parser.Constant{Value: 1},
		Rhs: parser.Variable{Value: "x"},
	}
	res := parser.AST{
		Lhs: parser.Subtraction{
			Lhs: parser.Constant{Value: 1},
			Rhs: parser.Variable{Value: "x"},
		},
		Rhs: parser.Constant{Value: 0},
	}

	assert.Equal(t, res, *transferTermRtoL(&ast, false))
}

func TestTransferFunction1(t *testing.T) {
	// 1 = x-5
	ast := parser.AST{
		Lhs: parser.Constant{Value: 1},
		Rhs: parser.Subtraction{Lhs: parser.Variable{Value: "x"}, Rhs: parser.Constant{Value: 5}},
	}

	res := parser.AST{
		Lhs: parser.Addition{Lhs: parser.Constant{Value: 1}, Rhs: parser.Constant{Value: 5}},
		Rhs: parser.Variable{Value: "x"},
	}

	assert.Equal(t, res, *transferTermRtoL(&ast, false))
}

func TestTransferFunction2(t *testing.T) {
	// 1 = x-5
	ast := parser.AST{
		Rhs: parser.Constant{Value: 1},
		Lhs: parser.Subtraction{Lhs: parser.Variable{Value: "x"}, Rhs: parser.Constant{Value: 5}},
	}

	res := parser.AST{Lhs: parser.Subtraction{Lhs: parser.Subtraction{Lhs: parser.Variable{Value: "x"}, Rhs: parser.Constant{Value: 5}}, Rhs: parser.Constant{Value: 1}}, Rhs: parser.Constant{Value: 0}}

	assert.Equal(t, res, *transferTermRtoL(&ast, false))
}

func TestTransferFunction3(t *testing.T) {

	// 1+y = x-5
	/// 1+y+5=x
	ast := parser.AST{
		Rhs: parser.Addition{Lhs: parser.Constant{Value: 1}, Rhs: parser.Variable{Value: "y"}},
		Lhs: parser.Subtraction{Lhs: parser.Variable{Value: "x"}, Rhs: parser.Constant{Value: 5}},
	}

	res := parser.AST{Lhs: parser.Subtraction{Lhs: parser.Subtraction{Lhs: parser.Variable{Value: "x"}, Rhs: parser.Constant{Value: 5}}, Rhs: parser.Constant{Value: 1}}, Rhs: parser.Constant{Value: 0}}

	assert.Equal(t, res, *transferTermRtoL(&ast, false))
}
