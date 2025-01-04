package rpn

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/al-zebra/utils"
)

type RPN struct {
	tokens []lexer.Token
}

func (rpn *RPN) Debug() {
	fmt.Println("RPN tokens: ", rpn.tokens)
}

func handleOperation(first, second *string, op func(int, int) int) (string, error) {
	// This will be nil if the stack is empty as first and second are meant to be poped off the stack in the main function
	if first == nil || second == nil {
		return "", errors.New("Stack is empty while calculating RPN-equation. This means that your equation is invalid or my RPNConverstion function is bugged.")
	}

	// Extraction, Main logic

	firstNumber, err := strconv.Atoi(*first)
	if err != nil {
		return "", fmt.Errorf("This error is not meant to be seen. As the error checking is already done above. Error:\n%s", err.Error())
	}

	secondNumber, err := strconv.Atoi(*second)
  
	if err != nil {
		return "", fmt.Errorf("This error is not meant to be seen. As the error checking is already done above. Error:\n%s", err.Error())
	}

	result := op(firstNumber, secondNumber)

	return strconv.Itoa(result), nil
}

func (tokens RPN) RPNCalc() ([]string, error) {
	var stack utils.Stack[string]
	for _, token := range tokens.tokens {
		switch token.Type {
		case lexer.DIVIDE:
			function := func(f, s int) int {
				return f / s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err
			}
			stack.PushBack(res)
		case lexer.MINUS:
			function := func(f, s int) int {
				return f - s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err
			}
			stack.PushBack(res)
		case lexer.MULTIPLY:
			function := func(f, s int) int {
				return f * s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err
			}
			stack.PushBack(res)
		case lexer.PLUS:
			function := func(f, s int) int {
				return f + s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err
			}
			stack.PushBack(res)
		case lexer.ROOT:
			function := func(f, s int) int {
				return int(math.Pow(float64(f), float64(s)))
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err
			}
			stack.PushBack(res)
		case lexer.VARIABLE:
			stack.PushBack(token.Value)
		case lexer.NUMBER:
			stack.PushBack(token.Value)
		}
	}
	return stack, nil
}

func helperNewOperation(lhs, rhs parser.Term, ty lexer.TokenType) parser.Term {
	switch ty {
	case lexer.DIVIDE:
		op := parser.Division{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.MINUS:
		op := parser.Subtraction{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.MULTIPLY:
		op := parser.Multiplication{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.PLUS:
		op := parser.Addition{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.ROOT:
		op := parser.Root{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	default:
		return nil
	}
}

type Node struct {
	val string
	left *Node
	right *Node
}

func (n *Node) helpDebug(level int, prefix string) {
	if n == nil {
		return
	}

	ident := strings.Repeat(" ", level)
	fmt.Printf("%s%s%s\n", ident, prefix, n.val)

	if n.left != nil {
		n.left.helpDebug(level + 1, "L: ")
	}

	if n.right != nil {
		n.right.helpDebug(level + 1, "R: ")
	}
}

func (n *Node) Debug() {
	n.helpDebug(0, "R: ")
}

func treeifyHelper(ty lexer.TokenType, lhs, rhs parser.Term) parser.Term {
  switch ty {
	case lexer.PLUS:
    t := parser.Addition{
    	lhs: lhs,
    	rhs: rhs,
    }
    return t
	case lexer.MINUS:
    t := parser.Subtraction{
    	lhs: lhs,
    	rhs: rhs,
    }
    return t
	case lexer.MULTIPLY:
    t := parser.Multiplication{
    	lhs: lhs,
    	rhs: rhs,
    }
    return t
  case lexer.DIVIDE:
    t := parser.Division{
    	lhs: lhs,
    	rhs: rhs,
    }
    return t
	case lexer.ROOT:
    t := parser.Root{
    	lhs: lhs,
    	rhs: rhs,
    }
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

func RPNConverstion(tokens []lexer.Token) RPN {
	var precedence map[lexer.TokenType]int = make(map[lexer.TokenType]int)
	precedence[lexer.PLUS] = 2
	precedence[lexer.MINUS] = 2
	precedence[lexer.MULTIPLY] = 3
	precedence[lexer.DIVIDE] = 3
	// FUTURE SELF. exponentiation has precedence of 4
	precedence[lexer.ROOT] = 4

	var result utils.Stack[lexer.Token]
	var operationStack utils.Stack[lexer.Token]
	for _, token := range tokens {
		//fmt.Println(operationStack, result, token)
		switch token.Type {
		case lexer.PLUS, lexer.MINUS:
			operationStack.PushBack(token)
		case lexer.MULTIPLY, lexer.DIVIDE:
			// i.e. addition or subtraction
			if precedence[*&operationStack.PeekBack().Type] < precedence[lexer.MULTIPLY] {
				operationStack.PushBack(token)
			} else {
				item := *operationStack.PopBack()
				result = append(result, item)
				operationStack.PushBack(token)
			}
		case lexer.ROOT:
			// i.e. addition, subtraction, multiplication, division
			if precedence[*&operationStack.PeekBack().Type] < precedence[lexer.ROOT] {
				operationStack.PushBack(token)
			} else {
				item := *operationStack.PopBack()
				result = append(result, item)
				operationStack.PushBack(token)
			}
		case lexer.LEFT_PAREN:
			operationStack.PushBack(token)
		case lexer.RIGHT_PAREN:
			for (*operationStack.PeekBack()).Type != lexer.LEFT_PAREN {
				item := *operationStack.PopBack()
				result.PushFront(item)
			}
			if operationStack.PeekBack().Type == lexer.LEFT_PAREN {
				operationStack.PopBack()
			}

		case lexer.NUMBER, lexer.VARIABLE:
			result.PushFront(token)
		}
		fmt.Println()
	}

	result = append(result, operationStack...)
	return RPN{
		tokens: result,
	}
}
