package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	equalsubtraction "github.com/al-zebra/solver/simplifier/equal_subtraction"
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
	ex := "1-1=10+x"
	lex := lexer.New(ex)
	par, err := parser.Parse(*lex)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
	par.Debug()

	pt := equalsubtraction.EqualSubtraction{}
  simp := pt.GetSimplified(par.Lhs)
  da := simp.(parser.Constant)
	fmt.Println(da)
}
