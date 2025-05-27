package simplifier

import (
	"github.com/al-zebra/parser"
	"github.com/al-zebra/solver/logger"
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
func transferTermRtoL(ast *parser.AST) *parser.AST {
	leafLhs, isVariableLhs := parser.IsLeafNode(ast.Lhs).(parser.Variable)
	leafRhs, isConstantRhs := parser.IsLeafNode(ast.Rhs).(parser.Constant)

	if isVariableLhs && isConstantRhs {
		// Any equation in form like "x = 2" will be turned to "x-2=0"
		return &parser.AST{
			Lhs: parser.Subtraction{
				Lhs: leafLhs,
				Rhs: leafRhs,
			},
			Rhs: parser.Constant{Value: 0},
		}
	}

	if lhs, rhs := parser.IsLeafNode(ast.Lhs), parser.IsLeafNode(ast.Rhs); lhs != nil && rhs != nil && lhs != rhs {
		// This is a logical error as if this brank is entered it means that, two constant are equal (in the equation) but not actually equal.
		// Thats why we are returning nil here as it cannot be processed.
		return nil
	}

	if parser.IsLeafNode(ast.Lhs) != nil {
		// rhs is required to be a operation due to the logic-gate above.
    rhsOp, err := ast.Rhs.(parser.Operation)
    if !err {
      // SAFTY: This branch of logic is already dealt above
      panic("This should not be raised")
    }
    // Rhs operation token type (lexer)
    rhsOptt := parser.TermToTokenType(rhsOp)
    opposite := parser.OppositeTokenType(rhsOptt)
    operation := parser.OperationBuilder(opposite, ast.Lhs, *rhsOp.GetRhs())
    return &parser.AST{
    	Lhs: operation,
    	Rhs: *rhsOp.GetLhs(),
    }
	}

	if parser.IsLeafNode(ast.Rhs) != nil {
		// lhs is required to be a operation due to the logic-gate above.
    lhsOp, err := ast.Lhs.(parser.Operation)
    if !err {
      // SAFTY: This branch of logic is already dealt above
      panic("This should not be raised")
    }
    // Lhs operation token type (lexer)
    lhsOptt := parser.TermToTokenType(lhsOp)
    opposite := parser.OppositeTokenType(lhsOptt)
    // x = 3*y 
    // x/y = 3
    operation := parser.OperationBuilder(opposite, ast.Lhs, *lhsOp.GetRhs())
    return &parser.AST{
    	Lhs: operation,
    	Rhs: *lhsOp.GetRhs(),
    }
	}
	return nil
}

func (_ BasicAlgebra) Solver(ast *parser.AST, logger *logger.Logger) {

}
