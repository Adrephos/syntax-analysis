package src

import (
	"fmt"
	"log"
	"k8s.io/apimachinery/pkg/util/sets"
)

func FollowCalc(g grammar, symbol string, r string) sets.String {
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
					if firstBeta.Has("ε") && n_terminal != symbol && n_terminal != r {
						followSet = followSet.Union(FollowCalc(g, n_terminal, r))
					}
					followSet = followSet.Union(firstBeta)
				}
			}
		}
	}
	return followSet.Delete("ε")
}

func Follow(g grammar, s string) sets.String {
	return FollowCalc(g, s, s)
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
