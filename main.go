package main

import (
	"errors"
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)
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
			return errors.New("In your equation, there are not exactly one equal sign") 
		}
	}
	if is_seen {
		return nil
	} else {

			return errors.New("In your equation, there are not exactly one equal sign")
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
					return errors.New("in your equation, there are more than one type of errors. (multi-variable equations to be added in future versions)") 
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
				return errors.New("In your equation, There is a root expression that does not preceed a number")
			}
			next_token := (*tokens)[idx+1]

			if next_token.Type != lexer.NUMBER {
				return errors.New("In your equation, There is a root expression that does not preceed a number")
			}
		}
	}
	return nil
}

func main() {
	ex := "2 3 +"
	lex := lexer.New(ex).TokenizeAll()
	toks := parser.RPNCalc(lex)
	fmt.Println("tokens",toks)
}

