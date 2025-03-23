package simplifier

import "github.com/al-zebra/parser"

type SimplificationPattern interface {
	IsValid(ast *parser.AST) bool
	Solver(ast *parser.AST, logger *Logger)
}

type BasicAlgebra struct{}

func hasVariable(term parser.Term) bool {
	stack := []parser.Term{term}
	for len(stack) > 0 {
		term := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		op, isOp := term.(parser.Operation)
		_, isVari := term.(parser.Variable)
		switch {
		case isOp:
			stack = append(stack, *op.GetLhs())
			stack = append(stack, *op.GetRhs())
		case isVari:
			return true
		}
	}
	return false
}

func (_ BasicAlgebra) IsValid(ast *parser.AST) bool {
	stack := []parser.Term{ast.Lhs, ast.Rhs}
	for len(stack) > 0 {
		term := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		op, isOp := term.(parser.Operation)
		switch {
		case isOp:
			if exp, isExp := op.(parser.Exponentiation); isExp {
				if hasVariable(exp) {
					return false
				}
			}
			stack = append(stack, *op.GetLhs())
			stack = append(stack, *op.GetRhs())
		}

	}
	return true
}
