package src

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
)

type grammar struct {
	N sets.String
	T sets.String
	P map[string][]string
	S string
}

func NewGrammar(N sets.String, T sets.String, P map[string][]string, S string) grammar {
	return grammar{
		N: N,
		T: T,
		P: P,
		S: S,
	}
}

func (g grammar) Print() {
	fmt.Println("Simbolo inicial: ", g.S)
	fmt.Println("Terminales: ", g.T.List())
	fmt.Println("No terminales: ", g.N.List())
	fmt.Println("Prducciones: ")
	for key, value := range g.P {
		fmt.Println(" ", key, "->", value)
	}

}

func derivationSet(g grammar) map[string]sets.String {
	symbolMap := make(map[string]sets.String)
	for symbol := range g.P {
		derivationSet := sets.NewString()
		for _, production := range g.P[symbol] {
			if g.N.Has(string(production[0])) {
				derivationSet.Insert(string(production[0]))
			}
		}
		symbolMap[symbol] = derivationSet

	}
	return symbolMap
}

func (g grammar) HasLeftRecursion() bool {
	for A := range g.N {
		visited := []string{A}
		visitedS := sets.NewString(A)
		i := 0
		for {
			B := visited[i]
			if productions, ok := g.P[B]; ok {
				for _, production := range productions {
					if production != "ε" && string(production[0]) == A {
						return true
					}
					if g.N.Has(string(production[0])) &&  !visitedS.Has(string(production[0])) {
						visited = append(visited, (string(production[0])))
						visitedS.Insert(string(production[0]))
					}
				}
				if len(visited)-1 <= i {
					break
				} else {
					i++
				}
			}
		}
	}
	return false
}
