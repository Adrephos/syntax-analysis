package src

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

func (g grammar) IsLL1() bool {
	isLL1 := true
	for n_terminal := range g.N {
		if !checkNonTerminal(g, n_terminal) {
			isLL1 = false
		}
	}
	return isLL1
}

func checkNonTerminal(g grammar, s string) bool {
	firstSet := make(map[string]sets.String)
	for _, production := range g.P[s] {
		firstSet[production] = FirstString(g, production)
	}

	for production_a, first_a := range firstSet {
		for production_b, first_b := range firstSet {
			if production_a != production_b {
				// At most one of cu and ,D can derive the empty string.
				if first_b.Has("ε") {
					intersection := first_a.Intersection(Follow(g, s)).Len()
					if intersection > 0 { return false }
				}
				//  For no terminal a do both a and ,O derive strings beginning with a.
				intersection := first_a.Intersection(first_b).Len()
				if intersection > 0 { return false }
			}
		}
	}

	return true
}
