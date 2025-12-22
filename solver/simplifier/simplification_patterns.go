package simplifier

import (
	"errors"

	"github.com/al-zebra/parser"
	"github.com/al-zebra/solver/logger"
)

func doesTermContainVariable(term *parser.Term) bool {
	if _, x := (*term).(parser.Constant); x {
		return true
	}
  op, isOp := (*term).(parser.Operation)
  if !isOp {
    panic("THIS IS STUPID")
  }
  // Check LHS


	return false
}

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
func transferBaseCase(ast *parser.AST) (*parser.AST, error) {
	leafVariableLhs, isVariableLhs := parser.IsLeafNode(ast.Lhs).(parser.Variable)
	leafConstantRhs, isConstantRhs := parser.IsLeafNode(ast.Rhs).(parser.Constant)
	leafConstantLhs, isConstantLhs := parser.IsLeafNode(ast.Lhs).(parser.Constant)
	leafVariableRhs, isVariableRhs := parser.IsLeafNode(ast.Rhs).(parser.Variable)

	if isVariableLhs && isConstantRhs {
		// Any equation in form like "x = 2" will be turned to "x-2=0"
		return &parser.AST{
			Lhs: parser.Subtraction{
				Lhs: leafVariableLhs,
				Rhs: leafConstantRhs,
			},
			Rhs: parser.Constant{Value: 0},
		}, nil
	}

	if isVariableRhs && isConstantLhs {
		// Any equation in form like "2=x" will be turned to "0=x-2"
		return &parser.AST{
			Rhs: parser.Constant{Value: 0},
			Lhs: parser.Subtraction{
				Rhs: leafConstantLhs,
				Lhs: leafVariableRhs,
			},
		}, nil
	}

	if lhs, rhs := parser.IsLeafNode(ast.Lhs), parser.IsLeafNode(ast.Rhs); lhs != nil && rhs != nil && lhs != rhs {
		// This is a logical error as if this brank is entered it means that, two constant are equal (in the equation) but not actually equal.
		// Thats why we are returning nil here as it cannot be processed.
		return nil, nil
	}

	return nil, errors.New("Base cases does not cover this ast.")
}

func TransLhsTerm(ast *parser.AST) *parser.AST {
	_, err := ast.Lhs.(parser.Operation)
	if !err {
		return nil
	}

	return nil
}

// Transfer the top-most rhs term to the opposite side.
func transferRhsTerm(ast *parser.AST) *parser.AST {
	if v, e := transferBaseCase(ast); e != nil {
		return v
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

		switch lhsOp.(type) {
		case parser.Addition, parser.Subtraction, parser.Exponentiation, parser.Root:
			return &parser.AST{
				Lhs: parser.Subtraction{
					Lhs: ast.Lhs,
					Rhs: ast.Rhs,
				},
				Rhs: parser.Constant{Value: 0},
			}
		case parser.Multiplication, parser.Division:
			return &parser.AST{
				Lhs: parser.Division{
					Lhs: ast.Lhs,
					Rhs: ast.Rhs,
				},
				Rhs: parser.Constant{Value: 1},
			}
		}
	}

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

// Transfer the top-most Lhs term to the opposite side.
func TransferLhsTerm(ast *parser.AST) *parser.AST {
	if v, e := transferBaseCase(ast); e != nil {
		return v
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

		switch lhsOp.(type) {
		case parser.Addition, parser.Subtraction, parser.Exponentiation, parser.Root:
			return &parser.AST{
				Lhs: parser.Subtraction{
					Lhs: ast.Lhs,
					Rhs: ast.Rhs,
				},
				Rhs: parser.Constant{Value: 0},
			}
		case parser.Multiplication, parser.Division:
			return &parser.AST{
				Lhs: parser.Division{
					Lhs: ast.Lhs,
					Rhs: ast.Rhs,
				},
				Rhs: parser.Constant{Value: 1},
			}
		}
	}

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

// Transfer the top-most rhs term to the opposite side.
func (_ BasicAlgebra) Solver(ast *parser.AST, logger *logger.Logger) {

}
