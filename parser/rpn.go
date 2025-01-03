package parser

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/utils"
)

type RPN struct {
	tokens []lexer.Token
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
			function := func (f, s int) int {
				return f / s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err 
			}
			stack.PushBack(res)
		case lexer.MINUS:
			function := func (f, s int) int {
				return f - s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err 
			}
			stack.PushBack(res)
		case lexer.MULTIPLY:
			function := func (f, s int) int {
				return f * s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err 
			}
			stack.PushBack(res)
		case lexer.PLUS:
			function := func (f, s int) int {
				return f + s
			}
			res, err := handleOperation(stack.PopBack(), stack.PopBack(), function)
			if err != nil {
				return []string{}, err 
			}
			stack.PushBack(res)
		case lexer.ROOT:
			function := func (f, s int) int {
				return int(math.Pow(float64(f), float64( s ))) 
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

func helperNewOperation(lhs, rhs Term, ty lexer.TokenType) Term {
	switch ty {
	case lexer.DIVIDE:
		op := Division{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.MINUS:
		op := Subtraction{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.MULTIPLY:
		op := Multiplication{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.PLUS:
		op := Addition{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	case lexer.ROOT:
		op := Root{
			lhs: lhs,
			rhs: rhs,
		}
		return op
	default:
		return nil
	}	
}

func (rpn RPN) Tree() Term {
	var stack utils.Stack[Term] 
	var m map[lexer.TokenType]bool
	m[lexer.DIVIDE] = true
	m[lexer.MINUS] = true
	m[lexer.PLUS] = true 
	m[lexer.MULTIPLY] = true
	m[lexer.ROOT] = true
	for _, tok := range rpn.tokens {
		_, isOp := m[tok.Type]
		if !isOp {
			ughh, err := strconv.Atoi(tok.Value)
			if err != nil {
				return nil
			}
			
			stack.PushFront(Constant{
				value: float32(ughh),
			})
		} else {
			left := stack.PopFront()
			right := stack.PopFront()	
			op := helperNewOperation(*left, *right, tok.Type)
			stack.PushFront(op)
		}
	}
	
	if len(stack) == 1 {
		return stack[0]
	} else {
		return nil
	}
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
		fmt.Println(operationStack, result, token)
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
