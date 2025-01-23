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

func GetPerformance() int {
	return 50
}

// x(x+2)
func GetSimplified(term parser.Term) parser.Term {
  op, isOp := term.(parser.Multiplication)
  if !isOp {
    return nil
  }
  leaf := parser.IsLeafNode(op.Lhs)
  if leaf != nil {
    if checkVariable(&op.Rhs) {
      
    }  
  }
}
