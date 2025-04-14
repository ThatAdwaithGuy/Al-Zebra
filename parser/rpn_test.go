package parser

import (
	"testing"

	"github.com/al-zebra/lexer"
)

func TestRpn(t *testing.T) {
  ex := "x^2+1=10"
  lx := lexer.New(ex)
  par, err := Parse(*lx)
  if err != nil {
    t.Fatalf("Parser thrown a error %s", err.Error())
  }
  par.Debug()
}
