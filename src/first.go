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

	if symbol == "" {
		firstSet.Insert(symbol)
	}

	if g.N.Has(symbol) {
		// Iterate over productions of X
		for _, production := range g.P[symbol] {
			if production == "" {
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
						if firstNonT.Has("") {
							firstSet = firstSet.Union(firstNonT.Delete(""))
						} else {
							firstSet = firstSet.Union(firstNonT)
							addEmpty = false
						}
					} else if g.N.Has(char) && !addEmpty {
						break
					} else if g.T.Has(char) && addEmpty{
						firstSet.Insert(char)
						addEmpty = false
					}
				}
				if addEmpty { firstSet.Insert("") }
			}
		}
	}
	return firstSet
}

func FirstSeveral(g grammar, str string) sets.String {
	firstSet := sets.NewString()
	addEmpty := true
	for i, runeChar := range str {
		char := string(runeChar)
		firstChar := First(g, char)
		if !firstChar.Has("") {
			firstSet = firstSet.Union(firstChar.Delete(""))
			addEmpty = false
			break
		}
		firstSet = firstSet.Union(firstChar.Delete(""))
		if i == len(str)-1 && addEmpty {
			firstSet.Insert("")
		}
	}
	return firstSet
}

func FirstGrammar(g grammar) {
	for _, value := range(g.N.List()) {
		fmt.Println(value, " -> ",First(g, value).List())
	}
}
