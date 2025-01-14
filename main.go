package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

type PointerTest struct {
	Hi  string
	Bye string
}

func (pt *PointerTest) GetHi() *string {
	return &pt.Hi
}

func (pt *PointerTest) GetBye() *string {
	return &pt.Bye
}

func main() {
	ex := "3x+1=10"
	lex := lexer.New(ex)
	par, err := parser.Parse(*lex)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
	par.Debug()

	var pt PointerTest
	*pt.GetBye() = "bye"
	*pt.GetHi() = "hi"

	ptt := PointerTest{
		Hi:  "hi",
		Bye: "bye",
	}
  fmt.Println("pt", pt)
  fmt.Println("ptt", ptt)
  fmt.Println("is equal", pt == ptt)
}
