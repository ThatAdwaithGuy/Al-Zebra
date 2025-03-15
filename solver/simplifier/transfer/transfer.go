package transfer

import "github.com/al-zebra/parser"

type TransferVariable struct {}

func (_ TransferVariable) GetPerformance() int {
  return 100
}
// Transfers the top most operation on LHS. if there is a constant term, lhs become 0 and// rhs changes.

func (_ TransferVariable) GetSimplifier(term parser.Term)  parser.Term {
  
}
