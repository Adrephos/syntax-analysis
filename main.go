package main

import (
	"fmt"
	"github.com/Adrephos/top-down/src"
)

func main() {
	grammar := src.GrammarInput()
	grammar.Print()
	fmt.Println("Recursividad por la izquierda: ", grammar.HasLeftRecursion())
	fmt.Printf("\nFirsts:\n")
	src.PrintFirstGrammar(grammar.FirstGrammar())
	fmt.Printf("\nFollows:\n")
	src.PrintFollowGrammar(grammar.FollowGrammar())
	fmt.Printf("\nEs LL(1): %v\n", grammar.IsLL1())
}
