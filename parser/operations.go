package parser

import (
	"errors"
	"math"

	"github.com/al-zebra/lexer"
)

// Just a bunch of errors
type UnhandledTermError struct{}

// Some helper function
// Returns nil if tt is not a operation
func OperationBuilder(tt lexer.TokenType, lhs, rhs Term) Term {
  switch tt {
  case lexer.DIVIDE:
	case lexer.EXPONENTIATION:
	case lexer.MINUS:
	case lexer.MULTIPLY:
	case lexer.PLUS:
	case lexer.ROOT:
	}
  return nil
}

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

// operations, like addition and subtraction should implement this interface
type Operation interface {
	Lhs() *Term
	Rhs() *Term
	Evaluate() (*Constant, error)
	IsTerm() bool
	getName() string
}

const DEPTH_LIMIT int = 100

// The main logic for evaluation of addition, subtraction, etc
func evaluateHelper(a Operation, operation func(float32, float32) float32, depth int) (*Constant, error) {
	// Just to make this perform better, I added this depth limit
	if depth > DEPTH_LIMIT {
		return nil, DepthError{}
	}
	// rhs and lhs values
	var lhsValue *float32
	var rhsValue *float32
	// Check if this a constant or not
	lhsConstantValue, lhsOk := (*a.Lhs()).(Constant)
	if lhsOk {
		lhsValue = &lhsConstantValue.Value
	} else {
		// It SHOULD be a operation as this function only supports term having term be as Constant or a Operation
		lhsOperationValue, lhsOperationOk := (*a.Lhs()).(Operation)
		if !lhsOperationOk {
			return nil, errors.New("UNHANDLED ERROR")
		}

		// the recursive part of the function
		lhsVal, err := evaluateHelper(lhsOperationValue, operation, depth+1)

		if err != nil {
			return nil, err
		}

		lhsValue = &lhsVal.Value
	}
	// The same logic as above but lhs is rhs now. To get more info, read the top part
	rhsConstantValue, rhsOk := (*a.Rhs()).(Constant)
	if rhsOk {
		rhsValue = &rhsConstantValue.Value
	} else {
		rhsOperationValue, rhsOperationOk := (*a.Rhs()).(Operation)
		if !rhsOperationOk {
			return nil, errors.New("UNHANDLED ERROR")

		}

		rhsVal, err := evaluateHelper(rhsOperationValue, operation, depth+1)
		if err != nil {
			return nil, err
		}

		rhsValue = &rhsVal.Value
	}

	// Why the heck go does not have PROPER null safty. I badly want a Option type
	if lhsValue != nil && rhsValue != nil {
		return &Constant{
			Value: operation(*lhsValue, *rhsValue),
		}, nil
	} else {
		return nil, UnhandledError{}
	}
}

// Operations

type Addition struct {
	lhs Term
	rhs Term
}


func (a Addition) Lhs() *Term {
	return &a.lhs
}

func (a Addition) Rhs() *Term {
	return &a.rhs
}

func (a Addition) IsTerm() bool {
	return true
}
func (a Addition) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 + f2
	}, 0)
}

func (a Addition) getName() string {
	return "+"
}

type Subtraction struct {
	lhs Term
	rhs Term
}


func (a Subtraction) Lhs() *Term {
	return &a.lhs
}

func (a Subtraction) Rhs() *Term {
	return &a.rhs
}

func (a Subtraction) IsTerm() bool {
	return true
}

func (a Subtraction) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 - f2
	}, 0)
}

func (a Subtraction) getName() string {
	return "-"
}

type Multiplication struct {
	lhs Term
	rhs Term
}


func (a Multiplication) Lhs() *Term {
	return &a.lhs
}

func (a Multiplication) Rhs() *Term {
	return &a.rhs
}

func (a Multiplication) IsTerm() bool {
	return true
}

func (a Multiplication) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 * f2
	}, 0)
}

func (a Multiplication) getName() string {
	return "*"
}

type Division struct {
	lhs Term
	rhs Term
}


func (a Division) Lhs() *Term {
	return &a.lhs
}

func (a Division) Rhs() *Term {
	return &a.rhs
}

func (a Division) IsTerm() bool {
	return true
}

func (a Division) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 / f2
	}, 0)
}

func (a Division) getName() string {
	return "/"
}

type Root struct {
	lhs Term
	rhs Term
}


func (a Root) Lhs() *Term {
	return &a.lhs
}

func (a Root) Rhs() *Term {
	return &a.rhs
}

func (a Root) IsTerm() bool {
	return true
}

func (a Root) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return float32(math.Pow(float64(f1), float64(1/f2)))
	}, 0)
}

func (a Root) getName() string {
	return "root"
}

type Exponentiation struct {
	lhs Term
	rhs Term
}


func (a Exponentiation) Lhs() *Term {
	return &a.lhs
}

func (a Exponentiation) Rhs() *Term {
	return &a.rhs
}

func (a Exponentiation) IsTerm() bool {
	return true
}

func (a Exponentiation) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return float32(math.Pow(
			float64(f1),
			float64(f2),
		))
	}, 0)
}

func (a Exponentiation) getName() string {
	return "^"
}
