package simplifier

import (
	"sync"

	"github.com/al-zebra/parser"
)

type Pattern interface {
  Matches(parser.Term) bool
  GetPerformance() int
  GetSimplified(parser.Term) parser.Term
}

type Simplifier struct {
  patterns []Pattern
  equation parser.Term 
}

// Creates a new simplifier.
func New(equation parser.Term) Simplifier {
  return Simplifier{
    patterns: []Pattern{},
    equation: equation,
  }
}

// Registers a pattern that can be used to simplify the equation.
func (s *Simplifier) Register(p Pattern) {
  s.patterns = append(s.patterns, p)
}

// Returns nil if there is no pattern that matches the equation.
func (s *Simplifier) Simplify() parser.Term { 
  pattern := s.matchPattern()
  if pattern != nil {
    return pattern.GetSimplified(s.equation)
  }
  return nil
}

// Gets the best pattern that matches the equation.
func (s *Simplifier) matchPattern() Pattern {
  var wg sync.WaitGroup
  
  allMatching := make( map[Pattern]int )
  for _, pattern := range s.patterns {
    wg.Add(1)
    go func(p Pattern) {
      defer wg.Done()

      if p.Matches(s.equation) {
        allMatching[p] = p.GetPerformance()
      }

    }(pattern)  
  }
  wg.Wait()
  
  var pattern Pattern
  maxValue := -1 * int(^uint(0) >> 1) 
  for p, perf := range allMatching {
    if perf > maxValue {
      pattern = p
      maxValue = perf
    }
  }

  return pattern
}
