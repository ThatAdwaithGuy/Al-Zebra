package solver

import (
	//	"errors"

	"errors"

	"github.com/al-zebra/parser"
)

func createOpposite(op parser.Operation, lhs, rhs parser.Term) parser.Operation {
	_, isAdd := op.(parser.Addition)
	_, isSub := op.(parser.Subtraction)
	_, isMul := op.(parser.Multiplication)
	_, isDiv := op.(parser.Division)
	_, isExp := op.(parser.Exponentiation)
	_, isRoo := op.(parser.Root)

	// Bozo time
	if isAdd {
		var ret parser.Subtraction
		*ret.Lhs() = lhs
		*ret.Rhs() = rhs
		return ret
	} else if isSub {
		var ret parser.Addition
		*ret.Lhs() = lhs
		*ret.Rhs() = rhs
		return ret
	} else if isMul {
		var ret parser.Division
		*ret.Lhs() = lhs
		*ret.Rhs() = rhs
		return ret
	} else if isDiv {
		var ret parser.Multiplication
		*ret.Lhs() = lhs
		*ret.Rhs() = rhs
		return ret
	} else if isExp {
		var ret parser.Root
		*ret.Lhs() = lhs
		*ret.Rhs() = rhs
		return ret
	} else if isRoo {
		var ret parser.Exponentiation
		*ret.Lhs() = lhs
		*ret.Rhs() = rhs
		return ret
	} else {
		return nil
	}
}

func CarryOperation(ast *parser.AST) (parser.AST, error) {
	lhsOp, isLhsOp := ast.Lhs.(parser.Operation)
	if !isLhsOp {
		return parser.AST{}, errors.New("Lhs is constant and cannot be carried")
	}
	getTop := lhsOp.Rhs()
	newRhs := createOpposite(lhsOp, ast.Rhs, *getTop)
  if newRhs == nil {
    return parser.AST{}, errors.New("Unhandled Error at solver.CarryOperation")
  }
	newLhs := *lhsOp.Lhs()
	retAST := parser.AST{
		Lhs: newLhs,
		Rhs: newRhs,
	}

	return retAST, nil
}
