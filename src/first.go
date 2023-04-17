package src

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
)

func First(g grammar, symbol string) sets.String {
	firstSet := sets.NewString()
	// If X ∈ Σ, then First(X) = {X}
	if g.T.Has(symbol) {
		firstSet.Insert(symbol)
	}
	if symbol == "ε" {
		firstSet.Insert(symbol)
	}
	if g.N.Has(symbol) {
		// Iterate over productions of X
		for _, production := range g.P[symbol] {
			if production == "ε" {
				firstSet.Insert(production)
				continue
			}
			//Check if first char is terminal
			firstChar := string(production[0])
			if g.T.Has(firstChar) {
				firstSet.Insert(firstChar)
			} else if g.N.Has(firstChar) {
				//If char is non-terminal iterate over all production
				addEmpty := true
				for _, value := range production {
					char := string(value)
					if g.N.Has(char) && addEmpty {
						firstNonT := First(g, char)
						if firstNonT.Has("ε") {
							firstSet = firstSet.Union(firstNonT.Delete("ε"))
						} else {
							firstSet = firstSet.Union(firstNonT)
							addEmpty = false
						}
					} else if g.N.Has(char) && !addEmpty {
						break
					} else if g.T.Has(char) && addEmpty {
						firstSet.Insert(char)
						addEmpty = false
					}
				}
				if addEmpty {
					firstSet.Insert("ε")
				}
			}
		}
	}
	return firstSet
}

func FirstString(g grammar, str string) sets.String {
	firstSet := sets.NewString()
	addEmpty := true
	if str == "ε" {
		firstSet.Insert("ε")
	}
	for i, runeChar := range str {
		char := string(runeChar)
		firstChar := First(g, char)
		if !firstChar.Has("ε") {
			firstSet = firstSet.Union(firstChar.Delete("ε"))
			addEmpty = false
			break
		}
		firstSet = firstSet.Union(firstChar.Delete("ε"))
		if i == len(str)-1 && addEmpty {
			firstSet.Insert("ε")
		}
	}
	return firstSet
}

func (g grammar) FirstGrammar() map[string]sets.String {
	firsts := make(map[string]sets.String)
	for _, value := range g.N.List() {
		firsts[value] = First(g, value)
	}
	for _, value := range g.T.List() {
		firsts[value] = First(g, value)
	}
	return firsts
}

func PrintFirstGrammar(fG map[string]sets.String) {
	for symbol, first := range fG {
		fmt.Println(symbol, "->", first.List())
	}
}
