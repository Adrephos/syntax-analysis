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

func NewGrammar( N sets.String, T sets.String, P map[string][]string, S string) grammar {
	return grammar{
		N: N,
		T: T,
		P: P,
		S: S,
	}
}

func (g grammar) Print() {
	fmt.Println("Simbolo inicial: ",g.S)
	fmt.Println("Terminales: ",g.T.List())
	fmt.Println("No terminales: ",g.N.List())
	fmt.Println("Prducciones: ")
	for key, value := range g.P {
		fmt.Println(" ", key, "->", value)
	}
	
}
