package parser

import (
	"github.com/al-zebra/lexer"
	"github.com/al-zebra/utils"
	"strconv"
)

func treeifyHelper(ty lexer.TokenType, lhs, rhs Term) Term {
	switch ty {
	case lexer.PLUS:
		t := NewAddition(lhs, rhs)
		return t
	case lexer.MINUS:
		t := NewSubtraction(lhs, rhs)
		return t
	case lexer.MULTIPLY:
		t := NewMultiplication(lhs, rhs)
		return t
	case lexer.DIVIDE:
		t := NewDivision(lhs, rhs)
		return t
	case lexer.ROOT:
		t := NewRoot(lhs, rhs)
		return t
	case lexer.EXPONENTIATION:
		t := NewExponentiation(lhs, rhs)
		return t
	default:
		return nil
	}
}

func Treeify(tokens RPN) *Term {
	var stack utils.Stack[Term]
	for _, tok := range tokens.tokens {
		switch tok.Type {
    case lexer.PLUS, lexer.MINUS, lexer.MULTIPLY, lexer.DIVIDE, lexer.ROOT, lexer.EXPONENTIATION:
			t := treeifyHelper(tok.Type, *stack.PopBack(), *stack.PopBack())
			stack.PushBack(t)
		case lexer.NUMBER:
			fl, err := strconv.ParseFloat(tok.Value, 32)
			if err != nil {
				return nil
			}
			t := Constant{
				Value: float32(fl),
			}
			stack.PushBack(t)

		case lexer.VARIABLE:
			t := Variable{
				Name: tok.Value,
			}
			stack.PushBack(t)
		}
	}
	return &stack[0]
}
