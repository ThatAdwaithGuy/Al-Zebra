package parser

import (

	"github.com/al-zebra/lexer"
)



// The main parser, has a Parser method which will generate a AST
type Parser struct {
	equation lexer.Lexer
	currAST AST
}

// As, 3x means 3 * x. this function will just expand 3x to 3 * x 
func MultiplyPass(tokens []lexer.Token) []lexer.Token {
	result := []lexer.Token{}

	for i := 0; i < len(tokens) - 1; i++ {
		currToken := tokens[i]
		nextToken := tokens[i + 1]

		if currToken.Type == lexer.NUMBER && nextToken.Type == lexer.VARIABLE {
			result = append(result, currToken)
			result = append(result, lexer.Token{
				Type: lexer.MULTIPLY,
				Value: "",
			})
		} else {
			result = append(result, currToken)
		}
	}

	return result 
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

// A constant number. like 1, 2, 1.2, 1.5
type Constant struct {
	value float32
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


