package parser

import (
  "math"
  "errors"
)

// Just a bunch of errors
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

// operations, like addition and subtraction should implement this interface
type Operation interface {
	Lhs() Term
	Rhs() Term
	Evaluate() (*Constant, error)
	IsTerm() bool
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
	lhsConstantValue, lhsOk := a.Lhs().(Constant)
	if lhsOk {
		lhsValue = &lhsConstantValue.Value
	} else {
		// It SHOULD be a operation as this function only supports term having term be as Constant or a Operation 
		lhsOperationValue, lhsOperationOk := a.Lhs().(Operation)
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
	rhsConstantValue, rhsOk := a.Rhs().(Constant)
	if rhsOk {
		rhsValue = &rhsConstantValue.Value
	} else {
		rhsOperationValue, rhsOperationOk := a.Lhs().(Operation)
		if !rhsOperationOk {
			return nil, errors.New("UNHANDLED ERROR")

		}
		
		rhsVal, err := evaluateHelper(rhsOperationValue, operation, depth+1)
		if err != nil {
			return nil, err
		}

		rhsValue = &rhsVal.Value
	}

	// Why the heck go does not have null safty. I badly want a Option type 
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

func NewAddition(lhs, rhs Term) Addition {
  return Addition{
  	lhs: lhs,
  	rhs: rhs,
  }
}

func (a Addition) Lhs() Term {
	return a.lhs
}

func (a Addition) Rhs() Term {
	return a.rhs
}

func (a Addition) IsTerm() bool {
	return true
}
func (a Addition) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 + f2
	}, 0)
}

type Subtraction struct {
	lhs Term
	rhs Term
}

func NewSubtraction(lhs, rhs Term) Subtraction {
  return Subtraction{
    lhs: lhs,
    rhs: rhs,
  }
}

func (a Subtraction) Lhs() Term {
	return a.lhs
}

func (a Subtraction) Rhs() Term {
	return a.rhs
}

func (a Subtraction) IsTerm() bool {
	return true
}
func (a Subtraction) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 - f2
	}, 0)
}

type Multiplication struct {
	lhs Term
	rhs Term
}

func NewMultiplication(lhs, rhs Term) Multiplication {
  return Multiplication{
    lhs: lhs,
    rhs: rhs,
  }
}

func (a Multiplication) Lhs() Term {
	return a.lhs
}

func (a Multiplication) Rhs() Term {
	return a.rhs
}

func (a Multiplication) IsTerm() bool {
	return true
}
func (a Multiplication) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 * f2
	}, 0)
}

type Division struct {
	lhs Term
	rhs Term
}

func NewDivision(lhs, rhs Term) Division{
  return Division{
    lhs: lhs,
    rhs: rhs,
  }
}

func (a Division) Lhs() Term {
	return a.lhs
}

func (a Division) Rhs() Term {
	return a.rhs
}

func (a Division) IsTerm() bool { 
  return true
}

func (a Division) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return f1 / f2
	}, 0)
}

type Root struct {
	lhs Term
	rhs Term
}

func NewRoot(lhs, rhs Term) Root {
  return Root{
      lhs: lhs,
      rhs: rhs,
  }
}



func (a Root) Lhs() Term {
	return a.lhs
}

func (a Root) Rhs() Term {
	return a.rhs
}

func (a Root) IsTerm() bool {
	return true
}

func (a Root) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return float32(math.Pow(float64(f1), float64(1/f2)))
	}, 0)
}
