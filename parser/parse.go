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
	IsTerm() bool
}

type Variable struct {
	Name string
}

func (v Variable) IsTerm() bool {
	return true
}

// A constant number. like 1, 2, 1.2, 1.5
type Constant struct {
	Value float32
}

func (c Constant) IsTerm() bool {
	return true
}
