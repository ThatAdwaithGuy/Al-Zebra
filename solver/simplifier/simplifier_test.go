package simplifier

import (
	"fmt"
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
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

type testID struct {
  id int
}

func (id testID) IsTerm() bool {
  return true
}

func (id testID) GetName() string {
  return fmt.Sprintf("ID: %d", id.id)
}



// TODO: Add more tests.



func TestSimplifySquential(t *testing.T) {
	patterns := []PatternMock{
		{
			id: 0,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Constant)
        if !isOp {
          return nil
        }
				return testID{
					id: 0,
				}
			},
			performance: 100,
		},
		{
			id: 1,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Operation)
        if !isOp {
          return nil
        }
				return testID{
					id: 1,
				}
			},
			performance: 10,
		},
		{
			id: 2,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Multiplication)
        if !isOp {
          return nil
        }
				return testID{
					id: 2,
				}
			},
			performance: 20,
		},
		{
			id: 3,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Addition)
        if !isOp {
          return nil
        }
				return testID{
					id: 3,
				} 
			},
			performance: 15,
		},
	}

  testCases := []parser.Term{
    parser.Constant{},
    parser.Root{},
    parser.Addition{},
    parser.Multiplication{},
  }
  var sim Simplifier
  for _, pattern := range patterns {
    sim.Register(pattern) 
  }
  for _, testCase := range testCases {
    sim.equation = testCase 
    pattern := sim.Simplify() 
    testid, isTestid := pattern.(testID)
    if !isTestid {
      t.Fatal("Pattern is not testID")
    }
    switch testCase {
    case parser.Constant{}:
      assert.Equal(t, testid.id, 0)
    case parser.Root{}:
      assert.Equal(t, testid.id, 1)
    case parser.Addition{}:
      assert.Equal(t, testid.id, 3)
    case parser.Multiplication{}:
      assert.Equal(t, testid.id, 2)
    } 
  }
}

func TestSimplify(t *testing.T) {
	patterns := []PatternMock{
		{
			id: 0,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Constant)
        if !isOp {
          return nil
        }
				return testID{
					id: 0,
				}
			},
			performance: 100,
		},
		{
			id: 1,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Operation)
        if !isOp {
          return nil
        }
				return testID{
					id: 1,
				}
			},
			performance: 10,
		},
		{
			id: 2,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Multiplication)
        if !isOp {
          return nil
        }
				return testID{
					id: 2,
				}
			},
			performance: 20,
		},
		{
			id: 3,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Addition)
        if !isOp {
          return nil
        }
				return testID{
					id: 3,
				} 
			},
			performance: 15,
		},
		{
			id: 4,
			matchFunc: func(term parser.Term) parser.Term {
				_, isOp := term.(parser.Subtraction)
        if !isOp {
          return nil
        }
				return testID{
					id: 4,
				}
			},
			performance: 14,
		},
	}

  testCases := []parser.Term{
    parser.Constant{},
    parser.Root{},
    parser.Addition{},
    parser.Subtraction{},
    parser.Multiplication{},
  }
  var sim Simplifier
  for _, pattern := range patterns {
    sim.Register(pattern) 
  }
  for _, testCase := range testCases {
    sim.equation = testCase 
    pattern := sim.Simplify() 
    testid, isTestid := pattern.(testID)
    if !isTestid {
      t.Fatal("Pattern is not testID")
    }
    switch testCase {
    case parser.Constant{}:
      assert.Equal(t, testid.id, 0)
    case parser.Root{}:
      assert.Equal(t, testid.id, 1)
    case parser.Addition{}:
      assert.Equal(t, testid.id, 3)
    case parser.Subtraction{}:
      assert.Equal(t, testid.id, 4)
    case parser.Multiplication{}:
      assert.Equal(t, testid.id, 2)
    } 
  }
}
