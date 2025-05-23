package main

import (

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

func main() {
	ex := "x*(x+1)=10"
	lex := lexer.New(ex)
	par, err := parser.Parse(*lex)
	if err != nil {
    return
	}
  par.Debug()
}
