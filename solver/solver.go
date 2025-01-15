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
    ret := parser.Subtraction{
      Lhs: lhs,
      Rhs: rhs,
    }
    return ret
	} else if isSub {
    ret := parser.Addition{
      Lhs: lhs,
      Rhs: rhs,
    }
    return ret
	} else if isMul {
    ret := parser.Division{
      Lhs: lhs,
      Rhs: rhs,
    }
    return ret
	} else if isDiv {
    ret := parser.Multiplication{
      Lhs: lhs,
      Rhs: rhs,
    }
    return ret
	} else if isExp {
    ret := parser.Root{
      Lhs: lhs,
      Rhs: rhs,
    }
    return ret
	} else if isRoo {
    ret := parser.Exponentiation{
      Lhs: lhs,
      Rhs: rhs,
    }
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
	getTop := lhsOp.GetRhs()
	newRhs := createOpposite(lhsOp, ast.Rhs, *getTop)
  if newRhs == nil {
    return parser.AST{}, errors.New("Unhandled Error at solver.CarryOperation")
  }
	newLhs := *lhsOp.GetLhs()
	retAST := parser.AST{
		Lhs: newLhs,
		Rhs: newRhs,
	}

	return retAST, nil
}
