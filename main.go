package main

import (
	"fmt"

	"github.com/al-zebra/lexer"
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
  fmt.Println( int(lexer.NUMBER) )
}

