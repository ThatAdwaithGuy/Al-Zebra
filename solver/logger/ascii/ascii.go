package ascii

import (
	"fmt"

	"github.com/al-zebra/parser"
)

type AsciiVisualization struct {
	ast *parser.AST
}

func New(ast *parser.AST) AsciiVisualization {
	return AsciiVisualization{
		ast: ast,
	}
}

func visualizeConstants(t parser.Term) (string, error) {
	switch c := t.(type) {
	case parser.Constant:
		return fmt.Sprintf("%g", c.Value), nil
	case parser.Variable:
		return c.Name, nil
	case parser.Operation:
		return "", fmt.Errorf("Passed in invalid type of parser")
	}
	return "", fmt.Errorf("Bad term")
}


func (v AsciiVisualization) Visualize() string {
	return ""
}
