package parser

import (
	"github.com/al-zebra/lexer"
)

// The main parser, has a Parser method which will generate a AST
type Parser struct {
	equation lexer.Lexer
	currAST AST
}

// Just a struct to hold the AST
type AST struct {
	lhs Term
	rhs Term
}

// A Term can be a constant or a unary method (like addition)
type Term interface {
	isTerm() bool
}

type Variable struct {
	Name string
}

func (v Variable) isTerm() bool {
	return true
}

// A constant number. like 1, 2, 1.2, 1.5
type Constant struct {
	Value float32
}

func (c Constant) isTerm() bool {
	return true
}
 
// operations, like addition and subtraction should implement this interface
type Operation interface {
	Lhs() Term
	Rhs() Term
	Evaluate() (*Constant, error)
	isTerm() bool
}

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


