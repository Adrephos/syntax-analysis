package main

import (
	"fmt"
	"github.com/Adrephos/top-down/src"
)

func main() {
	grammar := src.GrammarInput()
	fmt.Println(grammar.HasLeftRecursion())
	//fmt.Println("Firsts:")
	//src.PrintFirstGrammar(src.FirstGrammar(grammar))
	//fmt.Println("Follows:")
	//src.PrintFollowGrammar(src.FollowGrammar(grammar))
}
