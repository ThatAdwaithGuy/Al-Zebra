package parser

import (
	"fmt"
	"math"
	"strconv"

	"github.com/al-zebra/lexer"
)

type Slice[T any] []T

func (s *Slice[T]) pushBack(item T) {
	*s = append([]T{item}, *s...)
}

func (s *Slice[T]) pushFront(item T) {
	*s = append(*s, item)
}

func (s *Slice[T]) peekBack() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	last := (*s)[0]

	return &last
}

func (s *Slice[T]) peekFront() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	last := (*s)[length-1]

	return &last
}

func (s *Slice[T]) popBack() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	last := (*s)[0]
	*s = (*s)[1:]

	return &last
}

func (s *Slice[T]) popFront() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	last := (*s)[length-1]
	*s = (*s)[:length-2]

	return &last
}

// As, 3x means 3 * x. this function will just expand 3x to 3 * x
func MultiplyPass(tokens []lexer.Token) []lexer.Token {
	result := []lexer.Token{}

	for i := 0; i < len(tokens)-1; i++ {
		currToken := tokens[i]
		nextToken := tokens[i+1]

		if currToken.Type == lexer.NUMBER && nextToken.Type == lexer.VARIABLE {
			result = append(result, currToken)
			result = append(result, lexer.Token{
				Type:  lexer.MULTIPLY,
				Value: "",
			})
		} else {
			result = append(result, currToken)
		}
	}

	return result
}

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
	var stack Slice[string]
	for _, token := range tokens {
		fmt.Println(stack, token)
		switch token.Type {
		case lexer.DIVIDE:
			first := strcon(*stack.popBack())
			second := strcon(*stack.popBack())
			calc := first / second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.MINUS:
			first := strcon(*stack.popBack())
			second := strcon(*stack.popBack())
			calc := first - second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.MULTIPLY:
			first := strcon(*stack.popBack())
			second := strcon(*stack.popBack())
			calc := first * second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.PLUS:
			first := strcon(*stack.popBack())
			second := strcon(*stack.popBack())
			calc := first + second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.ROOT:
			first := strcon(*stack.popBack())
			second := strcon(*stack.popBack())
			calc := strconv.Itoa(int(math.Pow(float64(first), float64(second))))
			stack.pushBack(calc)
		case lexer.VARIABLE:
			stack.pushBack(token.Value)
		case lexer.NUMBER:
			stack.pushBack(token.Value)
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

	var result Slice[lexer.Token]
	var operationStack Slice[lexer.Token]
	for _, token := range tokens {
		fmt.Println(operationStack, result, token)
		switch token.Type {
		case lexer.PLUS, lexer.MINUS:
			operationStack.pushBack(token)
		case lexer.MULTIPLY, lexer.DIVIDE:
			// i.e. addition or subtraction
			if precedence[*&operationStack.peekBack().Type] < precedence[lexer.MULTIPLY] {
				operationStack.pushBack(token)
			} else {
				item := *operationStack.popBack()
				result = append(result, item)
				operationStack.pushBack(token)
			}
		case lexer.ROOT:
			// i.e. addition, subtraction, multiplication, division
			if precedence[*&operationStack.peekBack().Type] < precedence[lexer.ROOT] {
				operationStack.pushBack(token)
			} else {
				item := *operationStack.popBack()
				result = append(result, item)
				operationStack.pushBack(token)
			}
		case lexer.LEFT_PAREN:
			operationStack.pushBack(token)
		case lexer.RIGHT_PAREN:
			for (*operationStack.peekBack()).Type != lexer.LEFT_PAREN {
				item := *operationStack.popBack()
				result.pushFront(item)
			}
			if operationStack.peekBack().Type == lexer.LEFT_PAREN {
				operationStack.popBack()
			}

		case lexer.NUMBER, lexer.VARIABLE:
			result.pushFront(token)
		}
		fmt.Println()
	}
	result = append(result, operationStack...)
	return result
}
