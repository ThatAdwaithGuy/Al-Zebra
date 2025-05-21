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

  assert.Equal(t, true, hasVariable( lhs ))
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
