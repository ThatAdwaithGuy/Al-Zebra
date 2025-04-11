package ascii

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
)

func TestRPN(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	eq := "3x+1=10"
	lex := lexer.New(eq)
	toks := lex.TokenizeAll()
	hello, err := convertRPN(parser.RPN{Tokens: toks})
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("hello", hello)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Display the output in test log
	t.Log( output)
}
