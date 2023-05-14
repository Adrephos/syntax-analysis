package src

import (
	"fmt"
)

func getArgument(arguments map[argument]int, num int) argument {
	for argument, number := range arguments {
		if number == num {
			return argument
		}
	}
	return argument{}
}

func LRParsing(g grammar, s string, action table, arguments map[argument]int, err error) (bool, error) {
	// Parsing table M of G
	var stack StackInt
	// Conigure input as w$ where w is the string
	s = fmt.Sprintf("%s$", s)
	//start symbol S of G on top of the stack, above $
	stack.Push(0)

	index := 0

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	for {
		top := stack[len(stack)-1]
		a := string(s[index])
		actionValue := action.getValue(top, a)

		if actionValue.function == "s" {
			stack.Push(actionValue.number)
			index++
		} else if actionValue.function == "r" {
			arg := getArgument(arguments, actionValue.number)
			if arg.Production != "ε" {
				for i := 1; i <= len(arg.Production); i++ {
					stack.Pop()
				}
			}
			top = stack[len(stack)-1]
			stack.Push(action.getValue(top, arg.Symbol).number)
		} else if actionValue.function == "a" {
			break
		} else {
			return false, nil
		}
	}
	return true, nil
}
