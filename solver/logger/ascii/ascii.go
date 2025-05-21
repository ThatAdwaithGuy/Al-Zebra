package ascii

import (
	"fmt"

	"github.com/al-zebra/parser"
	"github.com/charmbracelet/bubbles/help"
)

type AsciiVisualization struct {
	ast *parser.AST
}

func New(ast *parser.AST) AsciiVisualization {
	return AsciiVisualization{
		ast: ast,
	}
}

func precedence(term parser.Term) int {
	switch expr.(type) {
	case parser.Addition, parser.Subtraction:
		return 1 
	case parser.Multiplication, parser.Division:
		return 2 
	case parser.Exponentiation, parser.Root:
		return 3
	default:
		return -1
	}
}


func helper(term parser.Term, prevPrec int) (string, error) {
  leaf_term := parser.IsLeafNode(term)
  if leaf_term != nil {
    switch v := leaf_term.(type) {
    case parser.Constant: 
      return fmt.Sprintf("%g", v .Value ), nil
    case parser.Variable: 
      return v.Value, nil
    }  
  }
  s := "" 
  currPrec = precedence(term)
  switch v := term.(type) {
  case parser.Addition:
    s = fmt.Sp
  case parser.Subtraction:
  case parser.Multiplication:
  case parser.Division:
  case parser.Exponentiation:
  case parser.Root:
  }

}

func (v AsciiVisualization) Visualize() (string, error) {
  return "", nil
}
