package main

import (
	"fmt"
	"github.com/Adrephos/top-down/src"
	"k8s.io/apimachinery/pkg/util/sets"
)

func main() {
	n := sets.NewString("S", "B", "C", "D", "E", "F")

	t := sets.NewString("a", "b", "c", "f", "g", "h")

	p := map[string][]string{
		"S": { "aBDh" },
		"B": { "cC" },
		"C": { "bc", "ε" },
		"D": { "EF" },
		"E": { "g", "ε" },
		"F": { "f", "ε" },
	}

	grammar := src.NewGrammar(n, t, p, "S")
	src.PrintFirstGrammar(src.FirstGrammar(grammar))
	fmt.Println(" ")

	n = sets.NewString("S", "A", "D", "B", "C")

	t = sets.NewString("a", "b", "d", "g")

	p = map[string][]string{
		"S": { "A" },
		"A": { "aBD" },
		"D": { "dD", "ε" },
		"B": { "b" },
		"C": { "g" },
	}
	
	grammar = src.NewGrammar(n, t, p, "S")
	src.PrintFirstGrammar(src.FirstGrammar(grammar))
	
	fmt.Println(" ")

	n = sets.NewString("E", "G", "T", "U", "F")

	t = sets.NewString("+", "*", "(", ")", "i")

	p = map[string][]string{
		"E": { "TG" },
		"G": { "+TG", "ε" },
		"T": { "FU" },
		"U": { "*FU", "ε" },
		"F": { "(E)", "i" },
	}
	
	grammar = src.NewGrammar(n, t, p, "E")
	src.PrintFirstGrammar(src.FirstGrammar(grammar))
	fmt.Println("Follows my friend")
	src.PrintFollowGrammar(src.FollowGrammar(grammar))
	fmt.Println(src.Follow(grammar, "E").List())
	fmt.Println(src.FirstString(grammar, "ε").List())	

	fmt.Println(" ")

	n = sets.NewString("A", "B", "C")

	t = sets.NewString("a", "b", "c")

	p = map[string][]string{
		"A": { "BC" },
		"B": { "ba", "ε" },
		"C": { "a", "ε" },
	}
	
	grammar = src.NewGrammar(n, t, p, "A")
	src.PrintFirstGrammar(src.FirstGrammar(grammar))
}
