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

func (s *Slice[T]) pop() *T {
	length := len(*s)
	if length == 0 {
		return nil
	}

	last := (*s)[0]
	*s = (*s)[1:]

	return &last
}

// As, 3x means 3 * x. this function will just expand 3x to 3 * x
func MultiplyPass(tokens []lexer.Token) []lexer.Token {
	result := []lexer.Token{}

	for i := 0; i < len(tokens) - 1; i++ {
		currToken := tokens[i]
		nextToken := tokens[i + 1]

		if currToken.Type == lexer.NUMBER && nextToken.Type == lexer.VARIABLE {
			result = append(result, currToken)
			result = append(result, lexer.Token{
				Type: lexer.MULTIPLY,
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
			first := strcon(*stack.pop())
			second := strcon(*stack.pop())
			calc := first / second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.MINUS:
			first := strcon(*stack.pop())
			second := strcon(*stack.pop())
			calc := first - second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.MULTIPLY:
			first := strcon(*stack.pop())
			second := strcon(*stack.pop())
			calc := first * second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.PLUS:
			first := strcon(*stack.pop())
			second := strcon(*stack.pop())
			calc := first + second
			stack.pushBack(strconv.Itoa(calc))
		case lexer.ROOT:
			first := strcon(*stack.pop())
			second := strcon(*stack.pop())
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
