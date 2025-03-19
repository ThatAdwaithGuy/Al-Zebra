package simplifier

import (
	"github.com/al-zebra/parser"
)

func TwoTermMerge(term parser.Term) (parser.Term, int) {
	op, isOp := term.(parser.Addition)

	if !isOp {
		return nil, 50
	}

	if op.Lhs == op.Rhs {
		return parser.Multiplication{
			Lhs: op.Lhs,
			Rhs: parser.Constant{Value: 2},
		}, 50
	} else {
		return nil, 50
	}
}

func constantOperation(term *parser.Term) bool {
	add, isAdd := (*term).(parser.Addition)
	if isAdd {
		return parser.IsLeafNode(add.Lhs) != nil && parser.IsLeafNode(add.Rhs) != nil
	}
	sub, isSub := (*term).(parser.Subtraction)
	if !isSub {
		return parser.IsLeafNode(sub.Lhs) != nil && parser.IsLeafNode(sub.Rhs) != nil
	}
	mul, isMul := (*term).(parser.Multiplication)
	if !isMul {
		return parser.IsLeafNode(mul.Lhs) != nil && parser.IsLeafNode(mul.Rhs) != nil
	}
	div, isDiv := (*term).(parser.Division)
	if !isDiv {
		return parser.IsLeafNode(div.Lhs) != nil && parser.IsLeafNode(div.Rhs) != nil
	}
	return false
}

// x(x+2)
func ExpandVariableConstantOperation(term parser.Term) (parser.Term, int) {
	op, isOp := term.(parser.Multiplication)
	if !isOp {
		return nil, 50
	}

	leaf := parser.IsLeafNode(op.Lhs)
	if leaf == nil {
		return nil, 50
	}

	if !constantOperation(&op.Rhs) {
		return nil, 50
	}

	opRhs, isOpRhs := op.Rhs.(parser.Operation)
	if !isOpRhs {
		return nil, 50
	}

	newLhs := parser.Multiplication{
		Lhs: op.Lhs,
		Rhs: *opRhs.GetLhs(),
	}
	newRhs := parser.Multiplication{
		Lhs: op.Lhs,
		Rhs: *opRhs.GetRhs(),
	}
	switch opRhs.(type) {
	case parser.Addition:
		return parser.Addition{
			Lhs: newLhs,
			Rhs: newRhs,
		}, 50
	case parser.Subtraction:
		return parser.Subtraction{
			Lhs: newLhs,
			Rhs: newRhs,
		}, 50
	case parser.Multiplication:
		return parser.Multiplication{
			Lhs: newLhs,
			Rhs: newRhs,
		}, 50
	case parser.Division:
		return parser.Division{
			Lhs: newLhs,
			Rhs: newRhs,
		}, 50
	case parser.Exponentiation:
		return parser.Exponentiation{
			Lhs: newLhs,
			Rhs: newRhs,
		}, 50
	case parser.Root:
		return parser.Root{
			Lhs: newLhs,
			Rhs: newRhs,
		}, 50
	}

	return nil, 50
}

// Justs simplifies if the top terms are equal.

func EqualSubtraction(term parser.Term) (parser.Term, int) {
	sub, isSub := term.(parser.Subtraction)
	if !isSub {
		return nil, 50
	}

	if sub.Lhs == sub.Rhs {
		return parser.Constant{
			Value: 0,
		}, 50
	} else {
		return nil, 50
	}
}

// Justs simplifies if the top terms are equal.
func ZeroOperation(term parser.Term) (parser.Term, int) {
	zero := parser.Constant{Value: 0}
	add, isAdd := term.(parser.Addition)
	sub, isSub := term.(parser.Subtraction)
	_, isMul := term.(parser.Multiplication)
	div, isDiv := term.(parser.Division)
	if isAdd {
		if add.Lhs == zero {
			return add.Rhs, 25
		} else if add.Rhs == zero {
			return add.Lhs, 25
		} else {
			return nil, 25
		}
	} else if isSub {
		if sub.Lhs == zero {
			return parser.Multiplication{
				Lhs: parser.Constant{
					Value: -1,
				},
				Rhs: sub.Rhs,
			}, 25
		} else if add.Rhs == zero {
			return add.Lhs, 25
		} else {
			return nil, 25
		}
	} else if isMul {
		return zero, 25
	} else if isDiv {
		if div.Rhs == zero {
			// Division by zero
			return nil, 25
		} else if div.Lhs == zero {
			return zero, 25
		} else {
			return nil, 25
		}
	}
	return nil, 25
}
