package equalsubtraction

import (
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

func TestEqualSubtractionMatch(t *testing.T) {
  term := parser.Subtraction{
  	Lhs: parser.Constant{Value: 1},
  	Rhs: parser.Constant{Value: 0},
  }

  e := EqualSubtraction{}
  assert.True(t, e.Matches(term))
  assert.Equal(t, parser.Constant{Value: 1}, e.GetSimplified(term))
}

func TestEqualSubtractionNoMatch(t *testing.T) {
  term := parser.Subtraction{
  	Lhs: parser.Constant{Value: 1},
  	Rhs: parser.Constant{Value: 4},
  }

  e := EqualSubtraction{}
  assert.False(t, e.Matches(term))
}
