package main

import (
	"fmt"
	"github.com/Adrephos/top-down/src"
)

func main() {
	grammar := src.GrammarInput()
	grammar.Print()
	fmt.Println(grammar.HasLeftRecursion())
	fmt.Printf("\nFirsts:\n")
	src.PrintFirstGrammar(grammar.FirstGrammar())
	/*
	fmt.Printf("\nFollows:\n")
	src.PrintFollowGrammar(grammar.FollowGrammar())
	*/
}
