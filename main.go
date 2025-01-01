package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

func main() {
	ex := "3 + 4 * 2"
	lex := lexer.New(ex).TokenizeAll()
	toks := parser.RPNConverstion(lex)
	stuff := parser.RPNCalc(toks)
	fmt.Println("tokens", toks, stuff)
}
