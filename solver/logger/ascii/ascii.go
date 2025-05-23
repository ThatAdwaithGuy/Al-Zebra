package ascii

import (
	"fmt"

	"github.com/al-zebra/parser"
)

type AsciiVisualization struct {
	ast *parser.AST
}

func New(ast *parser.AST) AsciiVisualization {
	return AsciiVisualization{
		ast: ast,
	}
}

func precedence(term parser.Term) int {
	switch term.(type) {
	case parser.Addition, parser.Subtraction:
		return 1
	case parser.Multiplication, parser.Division:
		return 2
	case parser.Exponentiation, parser.Root:
		return 3
	default:
		return -1
	}
}

func helper(term parser.Term, prevPrec int) string {
	leaf_term := parser.IsLeafNode(term)
	if leaf_term != nil {
		switch v := leaf_term.(type) {
		case parser.Constant:
			return fmt.Sprintf("%g", v.Value)
		case parser.Variable:
			return v.Value
		}
	}

	s := ""
	currPrec := precedence(term)

	switch v := term.(type) {
	case parser.Addition:
		s = fmt.Sprintf("%s + %s", helper(v.Lhs, currPrec), helper(v.Rhs, currPrec))
	case parser.Subtraction:
		s = fmt.Sprintf("%s - %s", helper(v.Lhs, currPrec), helper(v.Rhs, currPrec))
	case parser.Multiplication:
		s = fmt.Sprintf("%s * %s", helper(v.Lhs, currPrec), helper(v.Rhs, currPrec))
	case parser.Division:
		s = fmt.Sprintf("%s / %s", helper(v.Lhs, currPrec), helper(v.Rhs, currPrec))
	case parser.Exponentiation:
		s = fmt.Sprintf("%s ^ %s", helper(v.Lhs, currPrec), helper(v.Rhs, currPrec))
	case parser.Root:
		s = fmt.Sprintf("root%s(%s)", helper(v.Lhs, currPrec), helper(v.Rhs, currPrec))
	}

	if currPrec < prevPrec {
		return "(" + s + ")"
	}

	return s
}

func (v AsciiVisualization) Visualize() (string, error) {
	lhs := helper(v.ast.Lhs, -1)
	rhs := helper(v.ast.Rhs, -1)
	return lhs + " = " + rhs, nil
}
