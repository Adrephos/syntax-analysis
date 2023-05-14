package main

import (
	"fmt"

	"github.com/Adrephos/top-down/src"
	//"k8s.io/apimachinery/pkg/util/sets"
)

func main() {
	grammar := src.GrammarInput()

	stringArr := src.StringsInput()

	leftRec := false
	fmt.Printf("\nTop down parser\n\n")
	// LL1 parse routine
	if grammar.HasLeftRecursion() {
		fmt.Println("Error: La gramática tiene recursividad por la izquierda")
		leftRec = true
	}

	if grammar.IsLL1() && !leftRec {
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

	fmt.Printf("\n\nBottom up parser\n\n")
	// LR parse routine
	action, arguments, err := grammar.CreateSLRTable()
	for _, value := range stringArr {
			isPartOfGrammar, err := src.LRParsing(grammar, value, action, arguments, err)
			if isPartOfGrammar {
				fmt.Printf("La cadena %s SI hace parte de la gramática\n", value)
			} else {
				fmt.Printf("La cadena %s NO hace parte de la gramática\n", value)
				if err != nil {
					fmt.Printf(err.Error())
				}
			}
		}
}
