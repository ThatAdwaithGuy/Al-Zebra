package rpn
import (
  "strconv"
	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/al-zebra/utils"
	"github.com/al-zebra/operations"
)

func treeifyHelper(ty lexer.TokenType, lhs, rhs parser.Term) parser.Term {
  switch ty {
	case lexer.PLUS:
    t := operations.NewAddition(lhs, rhs)
    return t
	case lexer.MINUS:
    t := operations.NewSubtraction(lhs, rhs)
    return t
	case lexer.MULTIPLY:
    t := operations.NewMultiplication(lhs, rhs)
    return t
  case lexer.DIVIDE:
    t := operations.NewDivision(lhs, rhs)
    return t
	case lexer.ROOT:
    t := operations.NewRoot(lhs, rhs)
    return t
	default:
    return nil
	}
}

func Treeify(tokens RPN) *parser.Term {
	var stack utils.Stack[parser.Term]
	for _, tok := range tokens.tokens {
		switch tok.Type {
		case lexer.PLUS, lexer.MINUS, lexer.MULTIPLY, lexer.DIVIDE, lexer.ROOT:
      t := treeifyHelper(tok.Type, *stack.PopBack(), *stack.PopBack())
			stack.PushBack(t)
	  case lexer.NUMBER:
      fl, err :=strconv.ParseFloat(tok.Value, 32)
      if err != nil {
        return nil
      }
      t := parser.Constant{
        Value: float32(fl),
      }
      stack.PushBack(t)

	  case lexer.VARIABLE:
      t := parser.Variable{
        Name: tok.Value,
      }
      stack.PushBack(t)
	  }
	}
	return &stack[0] 
}

