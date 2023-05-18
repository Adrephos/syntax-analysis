package main

import (
	"os"

	"github.com/Adrephos/syntax-analysis/src"
)

func main() {

	argsWithoutProg := os.Args[1:]


	if len(argsWithoutProg) == 0 {

		grammar := src.GrammarInput()
		strings := src.StringsInput()

		grammar.Print()

		src.Ll1Routine(grammar, strings)
		src.Lr0Routine(grammar, strings)

	} else if argsWithoutProg[0] == "-g" {

		src.CreateGrammar()	

	} else {
		
		src.FileInput(argsWithoutProg[0])
	}
}
