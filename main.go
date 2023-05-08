package main

import (
	"fmt"
	"github.com/Adrephos/top-down/src"
)

func main() {
	grammar := src.GrammarInput()
	stringArr := src.StringsInput()
	if grammar.HasLeftRecursion() {
		fmt.Println("Error: La gramática tiene recursividad por la izquierda")
		return
	}

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
		fmt.Println("Error: Esta gramática no es LL(1)")
	}
}
