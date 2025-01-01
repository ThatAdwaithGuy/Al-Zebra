package parser

import (
	"fmt"
	"strconv"
	"math"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/utils"
)

func strcon(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	return result
}

// What is error handling
func RPNCalc(tokens []lexer.Token) []string {
	var stack utils.Stack[string]
	for _, token := range tokens {
		fmt.Println(stack, token)
		switch token.Type {
		case lexer.DIVIDE:
			first := strcon(*stack.PopBack())
			second := strcon(*stack.PopBack())
			calc := first / second
			stack.PushBack(strconv.Itoa(calc))
		case lexer.MINUS:
			first := strcon(*stack.PopBack())
			second := strcon(*stack.PopBack())
			calc := first - second
			stack.PushBack(strconv.Itoa(calc))
		case lexer.MULTIPLY:
			first := strcon(*stack.PopBack())
			second := strcon(*stack.PopBack())
			calc := first * second
			stack.PushBack(strconv.Itoa(calc))
		case lexer.PLUS:
			first := strcon(*stack.PopBack())
			second := strcon(*stack.PopBack())
			calc := first + second
			stack.PushBack(strconv.Itoa(calc))
		case lexer.ROOT:
			first := strcon(*stack.PopBack())
			second := strcon(*stack.PopBack())
			calc := strconv.Itoa(int(math.Pow(float64(first), float64(second))))
			stack.PushBack(calc)
		case lexer.VARIABLE:
			stack.PushBack(token.Value)
		case lexer.NUMBER:
			stack.PushBack(token.Value)
		}
	}
	return stack
}

func RPNConverstion(tokens []lexer.Token) []lexer.Token {
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
	return result
}
