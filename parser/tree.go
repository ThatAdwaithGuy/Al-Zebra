package parser

import (
	"fmt"
	"log"
	"strconv"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/utils"
)


func Treeify(tokens RPN) *Term {
	var stack utils.Stack[Term]
	for _, tok := range tokens.Tokens {

		switch tok.Type {
    case lexer.PLUS, lexer.MINUS, lexer.MULTIPLY, lexer.DIVIDE, lexer.ROOT, lexer.EXPONENTIATION:
      log.Println(tok, stack, len(stack))
      f := *stack.PopFront()
      s := *stack.PopFront()
      fmt.Println(f, s)
			t := OperationBuilder(tok.Type,s, f)
			stack.PushFront(t)
		case lexer.NUMBER:
			fl, err := strconv.ParseFloat(tok.Value, 32)
			if err != nil {
				return nil
			}
			t := Constant{
				Value: float32(fl),
			}
			stack.PushFront(t)

		case lexer.VARIABLE:
			t := Variable{
				Value: tok.Value,
			}
			stack.PushFront(t)
		}
	}
	return &stack[0]
}
