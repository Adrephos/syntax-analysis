package src

import (
	"fmt"
	"log"
	"k8s.io/apimachinery/pkg/util/sets"
)

func Follow(g grammar, symbol string) sets.String {
	if !g.N.Has(symbol) {
		log.Fatal(symbol, "is not a non-terminal symbol of this grammar")
	}
	followSet := sets.NewString()
	if g.S == symbol {
		followSet.Insert("$")
	}
	for n_terminal, productions := range g.P {
		for _, production := range productions {
			for i, char := range production {
				if string(char) == symbol {
					beta := "ε"
					if i != len(production)-1 {
						beta = production[i+1:]
					}
					firstBeta := FirstString(g, beta)
					if firstBeta.Has("ε") && n_terminal != symbol {
						followSet = followSet.Union(Follow(g, n_terminal))
					}
					followSet = followSet.Union(firstBeta)
				}
			}
		}
	}
	return followSet.Delete("ε")
}

func (g grammar) FollowGrammar() map[string]sets.String {
	follows := make(map[string]sets.String)
	for _, n_terminal := range g.N.List() {
		follows[n_terminal] = Follow(g, n_terminal)
		
	}
	return follows
}

func PrintFollowGrammar(fG map[string]sets.String) {
	for symbol, follow := range fG {
		fmt.Println(symbol, "->", follow.List())
	}
}
