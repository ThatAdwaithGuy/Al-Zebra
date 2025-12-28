package parser

import (
	"errors"
	"math"

	"github.com/al-zebra/lexer"
)

// Checks if term is either a constant or a variable, if it is then it return them 
// if not then returns nil.
func IsLeafNode(term Term) Term {
	c, isConstant := term.(Constant)
	v, isVariable := term.(Variable)

	if isConstant {
		return c
	} else if isVariable {
		return v
	} else {
		return nil
	}
}

func IsVariable(term Term) *Variable {
  c, isVariable := term.(Variable)
  if isVariable {
    return &c
  }
  return nil
}

func IsConstant(term Term) *Constant {
  c, isConstant := term.(Constant)
  if isConstant {
    return &c
  }
  return nil
}

func IsOperation(term Term) *Operation {
  c, isOperation := term.(Operation)
  if isOperation {
    return &c
  }
  return nil
}


// Just a bunch of errors
type UnhandledTermError struct{}

// Some helper function
// Return EOF if a non-operation term is given
func TermToTokenType(term Operation) lexer.TokenType {
	switch term.(type) {
	case Addition:
		return lexer.PLUS
	case Subtraction:
		return lexer.MINUS
	case Multiplication:
		return lexer.MULTIPLY
	case Division:
		return lexer.DIVIDE
	case Exponentiation:
		return lexer.EXPONENTIATION
	case Root:
		return lexer.ROOT
	}
	// SAFTY:
	// The switch statement is exhauted. meaning there is not other brach to brach off from.
	panic("if you see this, then a new operation is added but this function is not updated.")
}

// If this function returns EOF, that indicates a invalid token type given.
func OppositeTokenType(tt lexer.TokenType) lexer.TokenType {
	switch tt {
	case lexer.PLUS:
		return lexer.MINUS
	case lexer.MINUS:
		return lexer.PLUS
	case lexer.MULTIPLY:
		return lexer.DIVIDE
	case lexer.DIVIDE:
		return lexer.MULTIPLY
	case lexer.EXPONENTIATION:
		return lexer.ROOT
	case lexer.ROOT:
		return lexer.EXPONENTIATION
	}

	return lexer.EOF
}

// Returns nil if tt is not a operation
func OperationBuilder(tt lexer.TokenType, lhs, rhs Term) Term {
	switch tt {
	case lexer.PLUS:
		ret := Addition{
			Lhs: lhs,
			Rhs: rhs,
		}
		return ret
	case lexer.MINUS:
		ret := Subtraction{
			Lhs: lhs,
			Rhs: rhs,
		}
		return ret
	case lexer.MULTIPLY:
		ret := Multiplication{
			Lhs: lhs,
			Rhs: rhs,
		}
		return ret
	case lexer.DIVIDE:
		ret := Division{
			Lhs: lhs,
			Rhs: rhs,
		}
		return ret
	case lexer.EXPONENTIATION:
		ret := Exponentiation{
			Lhs: lhs,
			Rhs: rhs,
		}
		return ret
	case lexer.ROOT:
		ret := Root{
			Lhs: lhs,
			Rhs: rhs,
		}
		return ret
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
	GetLhs() *Term
	GetRhs() *Term
	Evaluate() (*Constant, error)
	IsTerm() bool
	GetName() string
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
	lhsConstantValue, lhsOk := (*a.GetLhs()).(Constant)
	if lhsOk {
		lhsValue = &lhsConstantValue.Value
	} else {
		// It SHOULD be a operation as this function only supports term having term be as Constant or a Operation
		lhsOperationValue, lhsOperationOk := (*a.GetLhs()).(Operation)
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
	rhsConstantValue, rhsOk := (*a.GetRhs()).(Constant)
	if rhsOk {
		rhsValue = &rhsConstantValue.Value
	} else {
		rhsOperationValue, rhsOperationOk := (*a.GetRhs()).(Operation)
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
	Lhs Term
	Rhs Term
}

func (a Addition) GetLhs() *Term {
	return &a.Lhs
}

func (a Addition) GetRhs() *Term {
	return &a.Rhs
}

func (a Addition) IsTerm() bool {
	return true
}

func (a Addition) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 + f2
	}, 0)
}

func (a Addition) GetName() string {
	return "+"
}

type Subtraction struct {
	Lhs Term
	Rhs Term
}

func (a Subtraction) GetLhs() *Term {
	return &a.Lhs
}

func (a Subtraction) GetRhs() *Term {
	return &a.Rhs
}

func (a Subtraction) IsTerm() bool {
	return true
}

func (a Subtraction) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 - f2
	}, 0)
}

func (a Subtraction) GetName() string {
	return "-"
}

type Multiplication struct {
	Lhs Term
	Rhs Term
}

func (a Multiplication) GetLhs() *Term {
	return &a.Lhs
}

func (a Multiplication) GetRhs() *Term {
	return &a.Rhs
}

func (a Multiplication) IsTerm() bool {
	return true
}

func (a Multiplication) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 * f2
	}, 0)
}

func (a Multiplication) GetName() string {
	return "*"
}

type Division struct {
	Lhs Term
	Rhs Term
}

func (a Division) GetLhs() *Term {
	return &a.Lhs
}

func (a Division) GetRhs() *Term {
	return &a.Rhs
}

func (a Division) IsTerm() bool {
	return true
}

func (a Division) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 / f2
	}, 0)
}

func (a Division) GetName() string {
	return "/"
}

type Root struct {
	Lhs Term
	Rhs Term
}

func (a Root) GetLhs() *Term {
	return &a.Lhs
}

func (a Root) GetRhs() *Term {
	return &a.Rhs
}

func (a Root) IsTerm() bool {
	return true
}

func (a Root) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return float32(math.Pow(float64(f1), float64(1/f2)))
	}, 0)
}

func (a Root) GetName() string {
	return "root"
}

type Exponentiation struct {
	Lhs Term
	Rhs Term
}

func (a Exponentiation) GetLhs() *Term {
	return &a.Lhs
}

func (a Exponentiation) GetRhs() *Term {
	return &a.Rhs
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

func (a Exponentiation) GetName() string {
	return "^"
}
