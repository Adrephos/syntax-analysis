package src

import (
	"errors"
	"fmt"
	"strings"
)

// Predictive top-down algorithm for LL1 gramamr
func PredictiveParsing(g grammar, s string) (bool, error) {
	// Parsing table M of G
	table, err := g.CreateTable()

	if err != nil {
		return false, err
	}

	var stack Stack
	// Conigure input as w$ where w is the string
	s = fmt.Sprintf("%s$", s)
	//start symbol S of G on top of the stack, above $
	stack.Push("$")
	stack.Push(g.S)
	top := stack[len(stack)-1]

	index := 0
	for top != "$" {
		a := string(s[index])
		top = stack[len(stack)-1]

		if top == a {
			stack.Pop()
			index++
		} else if g.T.Has(top) {
			return false, errors.New("Error de Stack, X es un terminal\n")
		} else if getValue(table, top, a) == "" {
			return false, errors.New(fmt.Sprintf("Error de Stack, M[%s, %s] es una celda de error\n", top, a))
		} else if len(getValue(table, top, a)) > 0 {
			production := strings.Split(getValue(table, top, a), "->")
			stack.Pop()
			if production[1] != "ε" {
				for i := len(production[1])-1; i >= 0; i-- {
					sy := production[1][i]
					syStr := string(sy)
					stack.Push(syStr)
				}
			}
		}
	}
	return true, nil
}
