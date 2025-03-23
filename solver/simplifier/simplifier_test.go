package simplifier

import (
	"fmt"
	"testing"

	"github.com/al-zebra/parser"
	"github.com/stretchr/testify/assert"
)

// Implment the Pattern interface
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
	patterns := []Pattern{
		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Constant)
			if !isOp {
				return nil, 100
			}
			return testID{
				id: 0,
			}, 100
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Operation)
			if !isOp {
				return nil, 10
			}
			return testID{
				id: 1,
			}, 10
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Multiplication)
			if !isOp {
				return nil, 20
			}
			return testID{
				id: 2,
			}, 20
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Addition)
			if !isOp {
				return nil, 15
			}
			return testID{
				id: 3,
			}, 15
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
	patterns := []Pattern{
		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Constant)
			if !isOp {
				return nil, 100
			}
			return testID{
				id: 0,
			}, 100
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Operation)
			if !isOp {
				return nil, 10
			}
			return testID{
				id: 1,
			}, 10
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Multiplication)
			if !isOp {
				return nil, 20
			}
			return testID{
				id: 2,
			}, 20
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Addition)
			if !isOp {
				return nil, 15
			}
			return testID{
				id: 3,
			}, 15
		},

		func(term parser.Term) (parser.Term, int) {
			_, isOp := term.(parser.Subtraction)
			if !isOp {
				return nil, 14
			}
			return testID{
				id: 4,
			}, 14
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
