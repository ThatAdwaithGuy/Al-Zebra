package termmerge

import "github.com/al-zebra/parser"


// A Pattern which is used two equal terms together into 
//      Mul
//      / \
//   Term  2
type TwoTermMerge struct{}

func (t TwoTermMerge) Matches(term parser.Term) bool {
  op, isOp := term.(parser.Addition) 
  if !isOp {
    return false
  }

  if op.Lhs == op.Rhs {
    return true
  }

  return false
}

func (t TwoTermMerge) GetPerformance() int                   { return 50 }

func (t TwoTermMerge) GetSimplified(term parser.Term) parser.Term {
  op, isOp := term.(parser.Addition)

  if !isOp {
    panic("Match should be run before simplifying")
  }

  return parser.Multiplication{
  	Lhs: op.Lhs,
  	Rhs: parser.Constant{Value: 2},
  }
}
