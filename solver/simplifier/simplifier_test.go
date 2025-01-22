package simplifier

import (
	"fmt"
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

type PatternMock struct {
	id          int
	matchFunc   func(parser.Term) parser.Term 
	performance int
}

// Implment the Pattern interface

func (p PatternMock) GetPerformance() int {
	return p.performance
}

// Do nothing here as we are not testing simplication logic.
// That testing will be done by each simplication module.
func (p PatternMock) GetSimplified(term parser.Term) parser.Term {
	return p.matchFunc(term) 
}


// TODO: Add more tests.
func TestSimplify(t *testing.T) {
	patterns := []PatternMock{
		{
			id: 0,
			matchFunc: func(term parser.Term) parser.Term {
				isOp,_ := term.(parser.Constant)
				return isOp
			},
			performance: 100,
		},
		{
			id: 1,
			matchFunc: func(term parser.Term) parser.Term {
				isOp,_ := term.(parser.Operation)
				return isOp
			},
			performance: 10,
		},
		{
			id: 2,
			matchFunc: func(term parser.Term) parser.Term {
				isOp,_ := term.(parser.Multiplication)
				return isOp
			},
			performance: 20,
		},
		{
			id: 3,
			matchFunc: func(term parser.Term) parser.Term {
				isAdd,_ := term.(parser.Addition)
				return isAdd
			},
			performance: 15,
		},
		{
			id: 4,
			matchFunc: func(term parser.Term) parser.Term {
				isAdd,_ := term.(parser.Subtraction)
				return isAdd
			},
			performance: 14,
		},
	}

	sim := New(nil)
	for _, pattern := range patterns {
		sim.Register(pattern)
	}
	testCases := []parser.Term{
		parser.Constant{},
		parser.Division{},
		parser.Multiplication{},
		parser.Addition{},
		parser.Subtraction{},
	}
	for _, term := range testCases {
		sim.equation = term
		pattern := sim.matchPattern()
		mock, isMock := pattern.(PatternMock)
		if !isMock {
			t.Fatal("Pattern is not a mock")
		}
		switch term {
		case parser.Constant{}:
			assert.Equal(t, 0, mock.id, "Simplification is not working for test case id 0")
		case parser.Division{}:
			assert.Equal(t, 1, mock.id, "Simplification is not working for test case id 1")
		case parser.Multiplication{}:
			assert.Equal(t, 2, mock.id, "Simplification is not working for test case id 2")
		case parser.Addition{}:
			assert.Equal(t, 3, mock.id, "Simplification is not working for test case id 3")
		case parser.Subtraction{}:
			assert.Equal(t, 4, mock.id, "Simplification is not working for test case id 4")
		}
	}
  
  for _, p := range patterns {
    fmt.Println(p.id) 
  }
}
