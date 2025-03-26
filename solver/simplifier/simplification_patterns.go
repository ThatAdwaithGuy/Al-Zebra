package simplifier

import (
	"fmt"

	"github.com/al-zebra/parser"
)

type SimplificationPattern interface {
	IsValid(ast *parser.AST) bool
	Solver(ast *parser.AST, fmtger *Logger)
}

type BasicAlgebra struct{}

func helper(t parser.Term, terms *[]parser.Term) {
	if t == nil {
		return
	}

	(*terms) = append((*terms), t)

	if op, isOp := t.(parser.Operation); isOp {
		(*terms) = append((*terms), op)
		helper(*op.GetLhs(), terms)
		helper(*op.GetRhs(), terms)
	}
}

func hasVariable(term parser.Term) bool {
	terms := []parser.Term{}

	helper(term, &terms)

	fmt.Println(terms)

	for _, t := range terms {
		if _, isVari := t.(parser.Variable); isVari {
			return true
		}
	}

	return false
}

func (_ BasicAlgebra) IsValid(ast *parser.AST) bool {
	stack := []parser.Term{ast.Lhs, ast.Rhs}
	for len(stack) > 0 {
		fmt.Println("before", stack)
		term := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(stack)
		if op, isOp := term.(parser.Operation); isOp {
			if exp, isExp := op.(parser.Exponentiation); isExp {
				if hasVariable(exp) {
					return false
				}
			}
			if exp, isExp := op.(parser.Root); isExp {
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
