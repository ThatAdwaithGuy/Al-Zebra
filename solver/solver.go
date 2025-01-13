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

	check := false
	slice := []bool{isAdd, isSub, isMul, isDiv, isExp, isRoo}
	for _, ele := range slice {
		if check && ele {
			return nil
		} else if !check && ele {
			check = true
		}
	}

	if isAdd {
		return parser.NewSubtraction(lhs, rhs)
	} else if isSub {
		return parser.NewAddition(lhs, rhs)
	} else if isMul {
		return parser.NewDivision(lhs, rhs)
	} else if isDiv {
		return parser.NewMultiplication(lhs, rhs)
	} else if isExp {
		return parser.NewRoot(lhs, rhs)
	} else if isRoo {
		return parser.NewExponentiation(lhs, rhs)
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
  newLhs := *lhsOp.Lhs()
  retAST := parser.AST{
  	Lhs: newLhs,
  	Rhs: newRhs,
  }
  
  return retAST, nil
}
