package ascii

import (
	"testing"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestConversion(t *testing.T)  {
  ex := "3x+1=10"
  lexer := lexer.New(ex)
  ast, err := parser.Parse(*lexer)
  if err != nil {
    return 
  }
  asc := AsciiVisualization{ast:ast}
  vis, err := asc.Visualize()
  if err != nil {
    return 
  }
  assert.Equal(t, "ad",  vis, "WRONG")
}
