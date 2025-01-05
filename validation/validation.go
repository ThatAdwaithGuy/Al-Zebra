package validation

import (
	"errors"
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/utils"
)

type Validation []lexer.Token

func (tokens *Validation) IsValid() error {
  onlyOne := tokens.OneTypeOfVariable()
  if onlyOne != nil {
    return onlyOne
  }
  
  onlyEqual := tokens.OnlyOneEqual()
  if onlyEqual != nil {
    return onlyEqual
  }

  root := tokens.RootPreceding()
  if root != nil {
    return root
  }

  return nil
}

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

type ErrorNoVariables []lexer.Token

func (e ErrorNoVariables) Error() string {
  return fmt.Sprintf("Your equation %s has no variables", e)
}


// OneTypeOfVariable Checks if there are only one type of variable in the equation
func (tokens *Validation) OneTypeOfVariable() error {
	var variableName *string
  variables := utils.Filter(*tokens, func(tok lexer.Token) bool {
    return tok.Type == lexer.VARIABLE
  })

  if len(variables) == 0 {
    return ErrorNoVariables(*tokens)
  }
  variableName = &variables[0].Value
  
  for _, ele := range variables {
    if ele.Value != *variableName {
      return errors.New("Equation has many variable types")
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


