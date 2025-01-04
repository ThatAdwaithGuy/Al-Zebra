package main

import (
	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser/rpn"
	"github.com/al-zebra/parser"
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

func main() {
	ex := "3 + 4 * 2"
	lex := lexer.New(ex).TokenizeAll()
	toks := rpn.RPNConverstion(lex)
	toks.Debug()
	//tree := parser.Treeify(toks)
}
