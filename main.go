package main

import (
	"fmt"
	"github.com/Adrephos/top-down/src"
)

func main() {
	grammar := src.GrammarInput()
	stringArr := src.StringsInput()

	if grammar.IsLL1() {
		for _, value := range stringArr {
			isPartOfGrammar, err := src.PredictiveParsing(grammar, value)
			if isPartOfGrammar {
				fmt.Printf("La cadena %s SI hace parte de la gramática\n", value)
			} else {
				fmt.Printf("La cadena %s NO hace parte de la gramática\n", value)
				fmt.Printf(err.Error())
			}
		}
	} else {
		fmt.Println("Esta gramática no es LL(1)")
	}
}
