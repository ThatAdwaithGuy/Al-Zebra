package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

func main() {
	ex := "x^2+1=10"
	lex := lexer.New(ex) 
  fmt.Println(lex.TokenizeAll())
  par, err := parser.Parse(*lex)
  if err != nil {
    fmt.Println("ERROR: ", err.Error())
  }
  par.Debug()  
}
