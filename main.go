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
		"C": { "bc", "" },
		"D": { "EF" },
		"E": { "g", "" },
		"F": { "f", "" },
	}

	grammar := src.NewGrammar(n, t, p, "S")
	src.FirstGrammar(grammar)
	fmt.Println(" ")

	n = sets.NewString("S", "A", "D", "B", "C")

	t = sets.NewString("a", "b", "d", "g")

	p = map[string][]string{
		"S": { "A" },
		"A": { "aBD" },
		"D": { "dD", "" },
		"B": { "b" },
		"C": { "g" },
	}
	
	grammar = src.NewGrammar(n, t, p, "S")
	src.FirstGrammar(grammar)
	
	fmt.Println(" ")

	n = sets.NewString("E", "G", "T", "U", "F")

	t = sets.NewString("+", "*", "(", ")", "i")

	p = map[string][]string{
		"E": { "TG" },
		"G": { "+TG", "" },
		"T": { "FU" },
		"U": { "*FU", "" },
		"F": { "(E)", "i" },
	}
	
	grammar = src.NewGrammar(n, t, p, "E")
	src.FirstGrammar(grammar)

	fmt.Println(" ")

	n = sets.NewString("A", "B", "C")

	t = sets.NewString("a", "b", "c")

	p = map[string][]string{
		"A": { "BC" },
		"B": { "ba", "" },
		"C": { "a", "" },
	}
	
	grammar = src.NewGrammar(n, t, p, "A")
	src.FirstGrammar(grammar)
	fmt.Println(src.FirstSeveral(grammar, "B"))	


}
