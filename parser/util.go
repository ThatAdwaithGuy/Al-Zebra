package parser

import (

	"github.com/al-zebra/lexer"
)

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

