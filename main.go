package main

import (
	// "fmt"

	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/al-zebra/solver/simplifier"
)

func main() {
	ex := "x*(x+1)=10"
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
