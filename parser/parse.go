package parser

import (
	"fmt"
	"github.com/al-zebra/lexer"
	"github.com/al-zebra/validation"
)

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

// Just a struct to hold the AST
type AST struct {
	Lhs Term
	Rhs Term
}

func Parse(lex lexer.Lexer) (*AST, error) {
	tokens := lex.TokenizeAll()

	// Validation
	valid := validation.Validation(tokens)
	isValid := valid.IsValid()
	if isValid != nil {
		return nil, isValid
	}
	tokens = MultiplyPass(tokens)

	var lhs []lexer.Token
	var rhs []lexer.Token
	isRhs := false

	for _, ele := range tokens {
		if ele.Type == lexer.EOF {
			break
		} else if ele.Type == lexer.EQUALS && !isRhs {
			isRhs = true
		} else if ele.Type != lexer.EQUALS && isRhs {
			rhs = append(rhs, ele)
		} else if ele.Type != lexer.EQUALS && !isRhs {
			lhs = append(lhs, ele)
		}
	}
	fmt.Println(lhs)
	lhsRPN := Converstion(lhs)
	//rhsRPN := Converstion(rhs)
	fmt.Println(lhsRPN)

	return nil, nil
}

// A Term can be a constant or a unary method (like addition)
type Term interface {
	IsTerm() bool
}

type Variable struct {
	Name string
}

func (v Variable) IsTerm() bool {
	return true
}

// A constant number. like 1, 2, 1.2, 1.5
type Constant struct {
	Value float32
}

func (c Constant) IsTerm() bool {
	return true
}
