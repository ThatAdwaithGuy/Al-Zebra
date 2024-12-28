package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdditionNoRecursion(t *testing.T) {
	addition := Addition{
		lhs: Constant{value: 1.0},
		rhs: Constant{value: 1.0},
	}
	value, err := addition.Evaluate()
	if err != nil {
		return
	}
	assert.Equal(t, &Constant{value: 2.0}, value, "AHADAHD")
}

func TestAdditionRecursion(t *testing.T) {
	oneAddition := Addition{
		lhs: Constant{value: 1.0},
		rhs: Constant{value: 1.0},
	}
	addition := Addition{
		lhs: oneAddition,
		rhs: Constant{value: 2.0},
	}

	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: 4.0}, value, "the result is not equal to 4")
}
