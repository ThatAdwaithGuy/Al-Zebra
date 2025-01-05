package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

func main() {
	//ex := "3x+1=10"
	new_ex := "3 - 2 + 1"
	lex := lexer.New(new_ex).TokenizeAll()
	lex = parser.MultiplyPass(lex)
	fmt.Println(lex)
	con := parser.Conversion(lex)

	//parse, err := parser.Parse(*lex)
	//if err != nil {
	//  fmt.Println("ERROR", err.Error())
	//}
	con.Debug()
}
