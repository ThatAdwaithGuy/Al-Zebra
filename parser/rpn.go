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

func (tokens RPN) Evaluate() ([]string, error) {
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

func helperNewOperation(lhs, rhs Term, ty lexer.TokenType) Term {
	switch ty {
	case lexer.DIVIDE:
		op := NewDivision(lhs, rhs)
		return op
	case lexer.MINUS:
		op := NewSubtraction(lhs, rhs)
		return op
	case lexer.MULTIPLY:
		op := NewMultiplication(lhs, rhs)
		return op
	case lexer.PLUS:
		op := NewAddition(lhs, rhs)
		return op
	case lexer.ROOT:
		op := NewRoot(lhs, rhs)
		return op
	default:
		return nil
	}
}
func Conversion(tokens []lexer.Token) RPN {
    // Initialize precedence map
    precedence := map[lexer.TokenType]int{
        lexer.PLUS:     2,
        lexer.MINUS:    2,
        lexer.MULTIPLY: 3,
        lexer.DIVIDE:   3,
        lexer.ROOT:     4,
    }

    var result []lexer.Token
    var operationStack []lexer.Token

    for _, token := range tokens {
        fmt.Println(token, operationStack, result)

        switch token.Type {
        case lexer.PLUS, lexer.MINUS, lexer.MULTIPLY, lexer.DIVIDE:
            for len(operationStack) > 0 {
                top := operationStack[len(operationStack)-1]
                if top.Type == lexer.LEFT_PAREN || precedence[top.Type] < precedence[token.Type] {
                    break
                }
                result = append(result, operationStack[len(operationStack)-1])
                operationStack = operationStack[:len(operationStack)-1]
            }
            operationStack = append(operationStack, token)

        case lexer.ROOT:
            for len(operationStack) > 0 {
                top := operationStack[len(operationStack)-1]
                if top.Type == lexer.LEFT_PAREN || precedence[top.Type] < precedence[lexer.ROOT] {
                    break
                }
                result = append(result, operationStack[len(operationStack)-1])
                operationStack = operationStack[:len(operationStack)-1]
            }
            operationStack = append(operationStack, token)

        case lexer.LEFT_PAREN:
            operationStack = append(operationStack, token)

        case lexer.RIGHT_PAREN:
            for len(operationStack) > 0 && operationStack[len(operationStack)-1].Type != lexer.LEFT_PAREN {
                result = append(result, operationStack[len(operationStack)-1])
                operationStack = operationStack[:len(operationStack)-1]
            }
            if len(operationStack) > 0 && operationStack[len(operationStack)-1].Type == lexer.LEFT_PAREN {
                operationStack = operationStack[:len(operationStack)-1] // Remove LEFT_PAREN
            }

        case lexer.NUMBER, lexer.VARIABLE:
            result = append(result, token)
        }

        fmt.Println()
    }

    // Pop remaining operators to result
    for len(operationStack) > 0 {
        result = append(result, operationStack[len(operationStack)-1])
        operationStack = operationStack[:len(operationStack)-1]
    }

    return RPN{
        tokens: result,
    }
}

