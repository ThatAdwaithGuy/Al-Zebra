package simplifier

import (
	"github.com/al-zebra/parser"
	"github.com/al-zebra/utils"
)

type Pattern interface {
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

func (s *Simplifier) matchPatternSequential() parser.Term {
	var bestPattern parser.Term
	bestPerformance := -1 * int(^uint(0)>>1)

	for _, pattern := range s.patterns {
		sim := pattern.GetSimplified(s.equation)
		if sim != nil {
			performance := pattern.GetPerformance()
			if performance > bestPerformance {
				bestPattern = sim
				bestPerformance = performance
			}
		}
	}

	return bestPattern
}

// Gets the best pattern that matches the equation and simplifies it.
func (s *Simplifier) Simplify() parser.Term {
	if len(s.patterns) < 4 {
		return s.matchPatternSequential()
	}

	results := make(chan utils.Tuple[parser.Term, int], len(s.patterns))
	for _, pattern := range s.patterns {
		go func(p Pattern) {
			sim := p.GetSimplified(s.equation)
			if sim != nil {
				results <- utils.Tuple[parser.Term, int]{
					F: sim,
					S: p.GetPerformance(),
				}
			} else {
				results <- utils.Tuple[parser.Term, int]{}
			}
		}(pattern)
	}

	var bestPattern parser.Term
	bestPerformance := -1 * int(^uint(0)>>1)

	for i := 0; i < len(s.patterns); i++ {
		result := <-results
		if result.F != nil && result.S > bestPerformance {
			bestPattern = result.F
			bestPerformance = result.S
		}
	}

	return bestPattern
}
