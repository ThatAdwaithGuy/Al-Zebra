package simplifier

import (
	"github.com/al-zebra/parser"
)

type SimplificationPattern interface {
	IsValid(*parser.AST) bool
	Solver(*parser.AST, *Logger)
}

type BasicAlgebra struct{}

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
	// TODO OPTIMIZATION: Can extract Lhs checker to stop the redundent Rhs check.
	return isValidHelper(ast.Lhs) && isValidHelper(ast.Rhs)
}



