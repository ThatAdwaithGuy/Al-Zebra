package parser

import "math"

type Addition struct {
	lhs Term
	rhs Term
}

func (a Addition) Lhs() Term {
	return a.lhs
}

func (a Addition) Rhs() Term {
	return a.rhs
}

func (a Addition) isTerm() bool {
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

func (a Subtraction) Lhs() Term {
	return a.lhs
}

func (a Subtraction) Rhs() Term {
	return a.rhs
}

func (a Subtraction) isTerm() bool {
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

func (a Multiplication) Lhs() Term {
	return a.lhs
}

func (a Multiplication) Rhs() Term {
	return a.rhs
}

func (a Multiplication) isTerm() bool {
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

func (a Division) Lhs() Term {
	return a.lhs
}

func (a Division) Rhs() Term {
	return a.rhs
}

func (a Division) isTerm() bool {
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

func (a Root) Lhs() Term {
	return a.lhs
}

func (a Root) Rhs() Term {
	return a.rhs
}

func (a Root) isTerm() bool {
	return true
}

func (a Root) Evaluate() (*Constant, error) {
	return evaluateHelper(a, func(f1, f2 float32) float32 {
		return float32(math.Pow(float64(f1), float64(1/f2)))
	}, 0)
}
