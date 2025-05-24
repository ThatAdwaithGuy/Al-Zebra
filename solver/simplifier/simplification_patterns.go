package simplifier

import (
	"github.com/al-zebra/parser"
	"github.com/al-zebra/solver/logger"
	"github.com/al-zebra/solver/logger/ascii"
)

// A interface to create a generic way to solve different "patterns" (forms of equations).
// IsValid method verifies if the equation is suitable (or possible) to solve (or in this case to apply that change)
// Solver method's answer will be logged into injected logger.
type SimplificationPattern interface {
	IsValid(*parser.AST) bool
	Solver(*parser.AST, *logger.Logger)
}

// BasicAlgebra pattern is to solve equations like:
// 3x+1=10
// 2x+5=10
// or a equaiton that does not exponentiation or root
type BasicAlgebra struct{}

// checks the equation (t) if it contains a varible
func helper(t parser.Term) bool {
	if t == nil {
		return false
	}
	if _, v := t.(parser.Variable); v {
		return true
	}

	if op, isOp := t.(parser.Operation); isOp {
		return helper(*op.GetLhs()) || helper(*op.GetRhs())
	}
	return false
}

func hasVariable(term parser.Term) bool {
	return helper(term)
}

// checks if there is a variable inside exponentiation or root operation
func isValidHelper(t parser.Term) bool {
	if t == nil {
		return true
	}

	if op, isOp := t.(parser.Operation); isOp {
		if exp, isExp := op.(parser.Exponentiation); isExp {
			return !hasVariable(exp)
		}
		if exp, isExp := op.(parser.Root); isExp {
			return !hasVariable(exp)
		}

		return isValidHelper(*op.GetLhs()) && isValidHelper(*op.GetRhs())

	}

	return true
}

func (_ BasicAlgebra) IsValid(ast *parser.AST) bool {
	return isValidHelper(ast.Lhs) && isValidHelper(ast.Rhs)
}

// Transfer the top-most term to the opposite side.
func transferTerm(ast *parser.AST) *parser.AST {
	if parser.IsLeafNode(ast.Lhs) != nil {
		// TODO
	}
	if parser.IsLeafNode(ast.Rhs) != nil {
		//TODO
	}
	return nil
}

func (_ BasicAlgebra) Solver(ast *parser.AST, logger *logger.Logger) {

}
