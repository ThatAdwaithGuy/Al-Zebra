package ascii

import (
	"fmt"
	"strings"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/al-zebra/utils"
)

type AsciiVisualization struct {
	ast *parser.AST
}

func New(ast *parser.AST) AsciiVisualization {
	return AsciiVisualization{
		ast: ast,
	}
}
func (v AsciiVisualization) Visualize() ( string ,error) {
  lhsVis := strings.Join(utils.Map( v.ast.LhsTokens , func(t lexer.Token) string { 
  return t.StringVisualization()
  }), "")

  rhsVis := strings.Join(utils.Map( v.ast.RhsTokens , func(t lexer.Token) string { 
  return t.StringVisualization()
  }), "")
  
  return fmt.Sprintf("%s = %s", lhsVis, rhsVis), nil
}
