package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

func main() {
	ex := "10=x*(x+1)"
	lex := lexer.New(ex)
	par, err := parser.Parse(*lex)
	if err != nil {
		return
	}
  
  fmt.Println(par.ToInfixNotation())
}
