package parser

import (
	"github.com/al-zebra/lexer"
)

// Constant is a term
// But all terms are not Constant
// 3x+1=10

type Parser struct {
	equation lexer.Lexer
	currAST AST
}


// As, 3x means 3 * x. this function will just expand 3x to 3 * x 
func multiplyPass(tokens []lexer.Token) []lexer.Token {
	for i := 1; i < len(tokens) - 1; i++ {
		
	}

	return []lexer.Token{}
}
type AST struct {
	lhs Term
	rhs Term
}

type Term interface {
	isTerm() bool
}

// This also should implement Term
type Constant struct {
	value float32
}

func (c Constant) isTerm() bool {
	return true
}

type Operation interface {
	Lhs() Term
	Rhs() Term
	Evaluate() (*Constant, error)
	isTerm() bool
}

type UnhandledTermError struct{}

func (e UnhandledTermError) Error() string {
	return "Got a unhandled term (Term which is not a operation nor a constant)"
}

type UnhandledError struct{}

func (e UnhandledError) Error() string {
	return "UNHANDLED ERROR"
}

type DepthError struct{}

func (e DepthError) Error() string {
	return "Depth of the equation's term execed the limit"
}

const DEPTH_LIMIT int = 10

func evaluateHelper(a Operation, operation func(float32, float32) float32, depth int) (*Constant, error) {
	if depth > DEPTH_LIMIT {
		return nil, DepthError{}
	}

	var lhsValue *float32 = nil
	var rhsValue *float32 = nil

	lhsConstantValue, lhsOk := a.Lhs().(Constant)
	if lhsOk {
		lhsValue = &lhsConstantValue.value
	} else {
		// It should be a operation
		lhsOperationValue, lhsOperationOk := a.Lhs().(Operation)
		if !lhsOperationOk {
			return nil, UnhandledTermError{}
		}
		// 	lhsVal, err := lhsOperationValue.Evaluate(depth + 1)
		lhsVal, err := evaluateHelper(lhsOperationValue, operation, depth+1)
		if err != nil {
			return nil, err
		}

		lhsValue = &lhsVal.value
	}

	rhsConstantValue, rhsOk := a.Rhs().(Constant)
	if rhsOk {
		rhsValue = &rhsConstantValue.value
	} else {
		// It should be a operation
		rhsOperationValue, rhsOperationOk := a.Lhs().(Operation)
		if !rhsOperationOk {
			return nil, UnhandledTermError{}
		}

		rhsVal, err := evaluateHelper(rhsOperationValue, operation, depth+1)
		if err != nil {
			return nil, err
		}

		rhsValue = &rhsVal.value
	}

	if lhsValue != nil && rhsValue != nil {
		return &Constant{
			value: operation(*lhsValue, *rhsValue),
		}, nil
	} else {
		return nil, UnhandledError{}
	}
}
