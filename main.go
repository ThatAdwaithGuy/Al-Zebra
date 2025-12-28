package main

import (
	// "fmt"

	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/al-zebra/solver/simplifier"
)

func main() {
	ex := "10=x*(x+1)"
	lex := lexer.New(ex)
	par, err := parser.Parse(*lex)
	if err != nil {
		return
	}

	fmt.Println("BEFORE")
	par.Debug()
	fmt.Println("AFTER")
	trans := simplifier.TransferLhsTerm(par)
	trans.Debug()
}
