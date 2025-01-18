package termmerge

import "github.com/al-zebra/parser"

type TermMerge struct{}

func (t TermMerge) Matches(term parser.Term) bool {
  op, isOp := term.(parser.Operation) 
  if !isOp {
    return false
  }

  
}
func (t TermMerge) GetPerformance() int                   { return 50 }
func (t TermMerge) GetSimplified(term parser.Term) parser.Term {}
