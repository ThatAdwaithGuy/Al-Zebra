package main

import (
	"github.com/al-zebra/lexer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualSignPass(t *testing.T) {
	test := "3x+1=10"
	l := lexer.New(test)
	tokens := Validation(l.TokenizeAll())
	check := tokens.OnlyOneEqual()

	assert.Nil(t, check, `"3x+1=10" equation only has 1 equal sign but the validation failed it`)
}

func TestEqualSignFail(t *testing.T) {
	test := "1=2=3" // Should fail
	l := lexer.New(test)
	tokens := Validation(l.TokenizeAll())
	check := tokens.OnlyOneEqual()

	assert.NotNil(t, check, `"1=2=3" equation does not have 1 equal sign but the validation passed it`)
}

func TestVariablePass(t *testing.T) {
	test := "3x+1=10"
	l := lexer.New(test)
	tokens := Validation(l.TokenizeAll())
	check := tokens.OneTypeOfVariable()

	assert.Nil(t, check, `"3x+1=10" equation only has 1 type of variable (x) but the validation failed it`)
}

func TestVariableFail(t *testing.T) {
	test := "3x=4y"
	l := lexer.New(test)
	tokens := Validation(l.TokenizeAll())
	check := tokens.OneTypeOfVariable()

	assert.NotNil(t, check, `"3x=4y" equation does not have 1 type of variable (has both x and y) but the validation passed it`)
}
/*
func TestRootPass(t *testing.T) {
	test := "3x+1=root2(10)"
	l := lexer.New(test)
	tokens := Validation(l.TokenizeAll())
	check := tokens.RootPreceeding()

	assert.Nil(t, check, `The root use here is correct but the validation function returned with error`)
}

func TestRootFail(t *testing.T) {
	test := "32+x=root(10)"
	l := lexer.New(test)
	tokens := Validation(l.TokenizeAll())
	check := tokens.RootPreceeding()

	assert.NotNil(t, check, `The root use here is incorrect but the validation function returned as a valid equation`)
}
*/

func TestCasesNormal() {
	var testCases map[string]int
	testCases["2+3"]

}
