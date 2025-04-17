package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/al-zebra/utils"
)
func main() {
  ex := "3x+1=10"
  lex := lexer.New(ex)
  ast, err := parser.Parse(*lex)
  if err != nil {
    return 
  }
  fmt.Println(utils.Map(ast.LhsTokens, func(t lexer.Token) string {
    return t.StringVisualization()
  }))
}

