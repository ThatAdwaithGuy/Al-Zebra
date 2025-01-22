package equalsubtraction

import (
	"github.com/al-zebra/parser"
)

// Justs simplifies if the top terms are equal.
type EqualSubtraction struct{}

func (e EqualSubtraction) GetPerformance() int {
	return 50
}

func (e EqualSubtraction) GetSimplified(term parser.Term) parser.Term {
	sub, isSub := term.(parser.Subtraction)
	if !isSub {
		return nil
	}

	if sub.Lhs == sub.Rhs {
		return parser.Constant{
			Value: 0,
		}
	} else {
		return nil
	}
}

// Justs simplifies if the top terms are equal.
type ZeroOperation struct{}

func (e ZeroOperation) GetPerformance() int {
	return 25
}

func (e ZeroOperation) GetSimplified(term parser.Term) parser.Term {
	zero := parser.Constant{Value: 0}
	add, isAdd := term.(parser.Addition)
	sub, isSub := term.(parser.Subtraction)
	_, isMul := term.(parser.Multiplication)
	div, isDiv := term.(parser.Division)
	if isAdd {
		if add.Lhs == zero {
			return add.Rhs
		} else if add.Rhs == zero {
			return add.Lhs
		} else {
			return nil
		}
	} else if isSub {
		return sub.Lhs
	} else if isMul {
		return zero
	} else if isDiv {
		if div.Rhs == zero {
			// Division by zero
			return nil
		} else if div.Lhs == zero {
			return zero
		} else {
			return nil
		}
	}
	return nil
}
