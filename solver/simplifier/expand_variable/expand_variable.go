package expandvariable

import "github.com/al-zebra/parser"

func checkVariable(term *parser.Term) bool {
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

// Java time
type ExpandVariableConstantOperation struct{}

func (_ ExpandVariableConstantOperation) GetPerformance() int {
	return 50
}

// x(x+2)
func (_ ExpandVariableConstantOperation) GetSimplified(term parser.Term) parser.Term {
	op, isOp := term.(parser.Multiplication)
	if !isOp {
		return nil
	}

	leaf := parser.IsLeafNode(op.Lhs)
	if leaf == nil {
		return nil
	}

	if !checkVariable(&op.Rhs) {
		return nil
	}

	opRhs, isOpRhs := op.Rhs.(parser.Operation)
	if !isOpRhs {
		return nil
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
		}
	case parser.Subtraction:
		return parser.Subtraction{
			Lhs: newLhs,
			Rhs: newRhs,
		}
	case parser.Multiplication:
		return parser.Multiplication{
			Lhs: newLhs,
			Rhs: newRhs,
		}
	case parser.Division:
		return parser.Division{
			Lhs: newLhs,
			Rhs: newRhs,
		}
	case parser.Exponentiation:
		return parser.Exponentiation{
			Lhs: newLhs,
			Rhs: newRhs,
		}
	case parser.Root:
		return parser.Root{
			Lhs: newLhs,
			Rhs: newRhs,
		}
	}

	return nil
}
