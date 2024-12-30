package main

import (
	"fmt"
	"math"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

type EqualSignErr bool

func (_ EqualSignErr) Error() string {
	return "In your equation, there are not exactly one equal sign"
}

type VariableErr bool

func (_ VariableErr) Error() string {
	return "In your equation, there are more than one type of errors. (Multi-variable equations to be added in future versions)"
}

type RootErr bool

func (_ RootErr) Error() string {
	return "In your equation, There is a root expression that does not preceed a number"
}

type Validation []lexer.Token

// Checks if there are more or less than one equal sign in the given equation
func (tokens *Validation) OnlyOneEqual() error {
	is_seen := false
	for _, tok := range *tokens {
		if tok.Type == lexer.EQUALS {
			if !is_seen {
				is_seen = true
				continue
			}
			return EqualSignErr(true)
		}
	}
	if is_seen {
		return nil
	} else {
		return EqualSignErr(true)
	}
}

// OneTypeOfVariable Checks if there are only one type of variable in the equation
func (tokens *Validation) OneTypeOfVariable() error {
	var variableName *string
	variableName = nil

	for _, tok := range *tokens {
		if tok.Type == lexer.VARIABLE {
			if variableName == nil {
				variableName = &tok.Value
			} else {
				if variableName != &tok.Value {
					return VariableErr(true)
				}
			}
		}
	}

	return nil
}

func (tokens *Validation) RootPreceding() error {
	for idx, tok := range *tokens {
		if tok.Type == lexer.ROOT {
			// out of bounds check
			if idx == len(*tokens)-1 {
				return RootErr(true)
			}
			next_token := (*tokens)[idx+1]

			if next_token.Type != lexer.NUMBER {
				return RootErr(true)
			}
		}
	}
	return nil
}

func main() {
	ex := "3x+(2*y)=10"
	lex := lexer.New(ex).TokenizeAll()
	toks := parser.MultiplyPass(lex)
	x := math.Pow(10, -1/2)
	fmt.Println(x)
	fmt.Println("tokens",toks)
}

