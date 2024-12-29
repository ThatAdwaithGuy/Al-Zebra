package parser

import (
	"math"
	"math/rand"
	"testing"

	"github.com/al-zebra/lexer"
	"github.com/stretchr/testify/assert"
)

// TODO: Make this test more concise and less verbose
// Disclosure: The random tests will sometimes fail as of floating point inaccuracy

func TestAdditionNoRecursion(t *testing.T) {
	addition := Addition{
		lhs: Constant{value: 1.0},
		rhs: Constant{value: 1.0},
	}
	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
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

func TestAdditionRandom(t *testing.T) {
	value1 := rand.Intn(100)
	value2 := rand.Intn(100)
	value3 := rand.Intn(100)

	result := (value1 + value2) + value3

	oneAddition := Addition{
		lhs: Constant{value: float32(value1)},
		rhs: Constant{value: float32(value2)},
	}
	addition := Addition{
		lhs: oneAddition,
		rhs: Constant{value: float32(value3)},
	}

	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: float32(result)}, value, "the result is not equal to 4")
}

func TestSubtractionNoRecursion(t *testing.T) {
	addition := Subtraction{
		lhs: Constant{value: 5.0},
		rhs: Constant{value: 3.0},
	}
	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}
	assert.Equal(t, &Constant{value: 2.0}, value, "AHADAHD")
}

func TestSubtractionRecursion(t *testing.T) {
	oneSubtraction := Subtraction{
		lhs: Constant{value: 5.0},
		rhs: Constant{value: 3.0},
	}
	addition := Subtraction{
		lhs: oneSubtraction,
		rhs: Constant{value: 3.0},
	}

	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: -1.0}, value, "the result is not equal to 4")
}

func TestSubtractionRandom(t *testing.T) {
	value1 := rand.Intn(100)
	value2 := rand.Intn(100)
	value3 := rand.Intn(100)

	result := (value1 - value2) - value3

	oneSubtraction := Subtraction{
		lhs: Constant{value: float32(value1)},
		rhs: Constant{value: float32(value2)},
	}
	subtraction := Subtraction{
		lhs: oneSubtraction,
		rhs: Constant{value: float32(value3)},
	}

	value, err := subtraction.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: float32(result)}, value, "the result is not equal to expected value")
}

func TestMultiplicationNoRecursion(t *testing.T) {
	addition := Multiplication{
		lhs: Constant{value: 10.0},
		rhs: Constant{value: 2.0},
	}
	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}
	assert.Equal(t, &Constant{value: 20.0}, value, "AHADAHD")
}

func TestMultiplicationRecursion(t *testing.T) {
	oneMultiplication := Multiplication{
		lhs: Constant{value: 10.0},
		rhs: Constant{value: 2.0},
	}
	addition := Multiplication{
		lhs: oneMultiplication,
		rhs: Constant{value: -1.0},
	}

	value, err := addition.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: -20.0}, value, "the result is not equal to 4")
}

func TestMultiplicationRandom(t *testing.T) {
	value1 := rand.Intn(100)
	value2 := rand.Intn(100)
	value3 := rand.Intn(100)

	result := (value1 * value2) * value3

	oneMultiplication := Multiplication{
		lhs: Constant{value: float32(value1)},
		rhs: Constant{value: float32(value2)},
	}
	multiplication := Multiplication{
		lhs: oneMultiplication,
		rhs: Constant{value: float32(value3)},
	}

	value, err := multiplication.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: float32(result)}, value, "the result is not equal to expected value")
}

func TestDivisionNoRecursion(t *testing.T) {
	division := Division{
		lhs: Constant{value: 10.0},
		rhs: Constant{value: 2.0},
	}
	value, err := division.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: 5.0}, value, "Expected 10/2 to equal 5")
}

func TestDivisionRecursion(t *testing.T) {
	oneDivision := Division{
		lhs: Constant{value: 20.0},
		rhs: Constant{value: 2.0},
	}
	division := Division{
		lhs: oneDivision,
		rhs: Constant{value: 2.0},
	}

	value, err := division.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: 5.0}, value, "Expected (20/2)/2 to equal 5")
}

func TestDivisionRandom(t *testing.T) {
	value1 := rand.Intn(100)
	value2 := rand.Intn(99) + 1 // Avoid division by zero
	value3 := rand.Intn(99) + 1 // Avoid division by zero

	result := (float32(value1) / float32(value2)) / float32(value3)

	oneDivision := Division{
		lhs: Constant{value: float32(value1)},
		rhs: Constant{value: float32(value2)},
	}
	division := Division{
		lhs: oneDivision,
		rhs: Constant{value: float32(value3)},
	}

	value, err := division.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: result}, value, "the result is not equal to expected value")
}

func TestRootNoRecursion(t *testing.T) {
	root := Root{
		lhs: Constant{value: 16.0},
		rhs: Constant{value: 2.0},
	}
	value, err := root.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}
	assert.Equal(t, &Constant{value: 4.0}, value, "Expected square root of 16 to equal 4")
}

func TestRootRecursion(t *testing.T) {
	oneRoot := Root{
		lhs: Constant{value: 16.0},
		rhs: Constant{value: 2.0},
	}
	root := Root{
		lhs: oneRoot,
		rhs: Constant{value: 2.0},
	}

	value, err := root.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: 2.0}, value, "Expected root of (root of 16) to equal 2")
}

func TestRootRandom(t *testing.T) {
	value1 := rand.Intn(100)
	value2 := rand.Intn(5) + 1 // Keep root index reasonable and non-zero
	value3 := rand.Intn(5) + 1 // Keep root index reasonable and non-zero

	result := float32(math.Pow(math.Pow(float64(value1), 1/float64(value2)), 1/float64(value3)))

	oneRoot := Root{
		lhs: Constant{value: float32(value1)},
		rhs: Constant{value: float32(value2)},
	}
	root := Root{
		lhs: oneRoot,
		rhs: Constant{value: float32(value3)},
	}

	value, err := root.Evaluate()
	if err != nil {
		t.Errorf("Got error from evaluate %s", err.Error())
	}

	assert.Equal(t, &Constant{value: result}, value, "the result is not equal to expected value")
}

func TestMultiplyPass(t *testing.T) {
	// Mock data, this is not valid
	ex := "3x+2x"
	tokens := lexer.New(ex).TokenizeAll()
	pass := MultiplyPass(tokens)
	expected := []lexer.Token{
		{
			Type:  lexer.NUMBER,
			Value: "3",
		},
		{
			Type: lexer.MULTIPLY,
			Value: "",
		},
		{
			Type: lexer.VARIABLE,
			Value: "x",
		},
		{
			Type: lexer.PLUS,
			Value: "",
		},
		{
			Type: lexer.NUMBER,
			Value: "2",
		},
		{
			Type: lexer.MULTIPLY,
			Value: "",
		},
		{
			Type: lexer.VARIABLE,
			Value: "x",
		},
	}
	assert.Equal(t, expected, pass, "Multiply pass results in wrong values")
}
