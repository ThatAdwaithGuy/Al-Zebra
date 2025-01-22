package equalsubtraction

import (
	"testing"

	"github.com/al-zebra/parser"
)

func TestEqualSubtraction_GetPerformance(t *testing.T) {
	es := EqualSubtraction{}
	if got := es.GetPerformance(); got != 50 {
		t.Errorf("GetPerformance() = %v, want %v", got, 50)
	}
}

func TestEqualSubtraction_GetSimplified(t *testing.T) {
	es := EqualSubtraction{}
	tests := []struct {
		name     string
		input    parser.Term
		expected parser.Term
	}{
		{
			name: "Equal terms subtraction",
			input: parser.Subtraction{
				Lhs: parser.Constant{Value: 5},
				Rhs: parser.Constant{Value: 5},
			},
			expected: parser.Constant{Value: 0},
		},
		{
			name: "Different terms subtraction",
			input: parser.Subtraction{
				Lhs: parser.Constant{Value: 5},
				Rhs: parser.Constant{Value: 3},
			},
			expected: nil,
		},
		{
			name:     "Non-subtraction operation",
			input:    parser.Constant{Value: 5},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := es.GetSimplified(tt.input)
			if got != tt.expected {
				t.Errorf("GetSimplified() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestZeroOperation_GetPerformance(t *testing.T) {
	zo := ZeroOperation{}
	if got := zo.GetPerformance(); got != 25 {
		t.Errorf("GetPerformance() = %v, want %v", got, 25)
	}
}

func TestZeroOperation_GetSimplified(t *testing.T) {
	zo := ZeroOperation{}
	zero := parser.Constant{Value: 0}
	five := parser.Constant{Value: 5}

	tests := []struct {
		name     string
		input    parser.Term
		expected parser.Term
	}{
		{
			name: "Addition with left zero",
			input: parser.Addition{
				Lhs: zero,
				Rhs: five,
			},
			expected: five,
		},
		{
			name: "Addition with right zero",
			input: parser.Addition{
				Lhs: five,
				Rhs: zero,
			},
			expected: five,
		},
		{
			name: "Subtraction with zero",
			input: parser.Subtraction{
				Lhs: five,
				Rhs: zero,
			},
			expected: five,
		},
		{
			name: "Multiplication with zero",
			input: parser.Multiplication{
				Lhs: five,
				Rhs: zero,
			},
			expected: zero,
		},
		{
			name: "Division with zero numerator",
			input: parser.Division{
				Lhs: zero,
				Rhs: five,
			},
			expected: zero,
		},
		{
			name: "Division by zero",
			input: parser.Division{
				Lhs: five,
				Rhs: zero,
			},
			expected: nil,
		},
		{
			name:     "Non-operation term",
			input:    parser.Constant{Value: 5},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zo.GetSimplified(tt.input)
			if got != tt.expected {
				t.Errorf("GetSimplified() = %v, want %v", got, tt.expected)
			}
		})
	}
}
