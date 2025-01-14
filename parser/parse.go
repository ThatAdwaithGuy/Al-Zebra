package parser

import (
	"fmt"
	"strings"

	"github.com/al-zebra/lexer"
	"github.com/al-zebra/validation"
)

func MultiplyPass(tokens []lexer.Token) []lexer.Token {
	result := []lexer.Token{}

	for i := 0; i < len(tokens)-1; i++ {
		currToken := tokens[i]
		nextToken := tokens[i+1]

		if currToken.Type == lexer.NUMBER && nextToken.Type == lexer.VARIABLE {
			result = append(result, currToken)
			result = append(result, lexer.Token{
				Type:  lexer.MULTIPLY,
				Value: "",
			})
		} else {
			result = append(result, currToken)
		}
	}

	return result
}

// Just a struct to hold the AST
type AST struct {
	Lhs Term
	Rhs Term
}

func helperDebug(term Term, level int) {
	if term == nil {
		return
	}

	indent := strings.Repeat(" ", level)
	fmt.Printf("%s%s\n", indent, term.getName())
	if op, isOp := term.(Operation); isOp {
		helperDebug(*op.Lhs(), level+1)
		helperDebug(*op.Rhs(), level+1)
	}
}

func (ast *AST) Debug() {
	fmt.Println("Lhs")
	helperDebug(ast.Lhs, 0)
	fmt.Println("")
	fmt.Println("Rhs")
	helperDebug(ast.Rhs, 0)
}

const DEBUG = true

func Parse(lex lexer.Lexer) (*AST, error) {
	tokens := lex.TokenizeAll()

	// Validation
	valid := validation.Validation(tokens)
	isValid := valid.IsValid()
	if isValid != nil {
		return nil, isValid
	}

	tokens = MultiplyPass(tokens)
	var lhs []lexer.Token
	var rhs []lexer.Token
	isRhs := false

	for _, ele := range tokens {
		if ele.Type == lexer.EOF {
			break
		} else if ele.Type == lexer.EQUALS && !isRhs {
			isRhs = true
		} else if ele.Type != lexer.EQUALS && isRhs {
			rhs = append(rhs, ele)
		} else if ele.Type != lexer.EQUALS && !isRhs {
			lhs = append(lhs, ele)
		}
	}

	lhsRPN := Conversion(lhs)
	rhsRPN := Conversion(rhs)

	lhsTree := Treeify(lhsRPN)
	rhsTree := Treeify(rhsRPN)

	ast := AST{
		Lhs: *lhsTree,
		Rhs: *rhsTree,
	}

	return &ast, nil
}

// A Term can be a constant or a unary method (like addition)
type Term interface {
	IsTerm() bool
	getName() string
}

type Variable struct {
	Name string
}

func (v Variable) IsTerm() bool {
	return true
}

func (v Variable) getName() string {
	return v.Name
}

// A constant number. like 1, 2, 1.2, 1.5
type Constant struct {
	Value float32
}

func (c Constant) IsTerm() bool {
	return true
}

func (c Constant) getName() string {
	return fmt.Sprintf("%f", c.Value)
}
