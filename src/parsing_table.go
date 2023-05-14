package src

import (
	"fmt"
	"log"
)

// Creates a parsing table for a LL1 grammar
func (g grammar) CreateTable() map[string]map[string]string {
	table := make(map[string]map[string]string)
	followGrammar := g.FollowGrammar()
	for symbol, productions := range g.P {
		//For each production A -> α of the grammar
		for _, production := range productions {
			firstOfProduction := FirstString(g, production)
			//For each terminal a ∈ firstOfProduction, add A -> production to M[A, a]
			for a := range firstOfProduction {
				if g.T.Has(a) {
					prod := fmt.Sprintf("%s->%s", symbol, production)
					if getValue(table, symbol, a) != "" {
						log.Fatal("Gramatica ambigua, fallo al crear la tabla")
					}
					addCell(table, symbol, a, prod)
				}
			}
			//If ε ∈ firstOfProduction, then for each terminal b or $ ∈ Follow(A), add A -> production to M[A, b]
			if firstOfProduction.Has("ε") {
				for b := range followGrammar[symbol] {
					if g.T.Has(b) || b == "$" {
						prod := fmt.Sprintf("%s->%s", symbol, production)
						if getValue(table, symbol, b) != "" {
							log.Fatal("Error: Gramatica ambigua, fallo al crear la tabla")
						}
						addCell(table, symbol, b, prod)
					}
				}
			}
		}
	}
	return table
}

func addCell(table map[string]map[string]string, row, column, value string) {
	if table[row] == nil {
		table[row] = make(map[string]string)
	}
	table[row][column] = value
}

func getValue(table map[string]map[string]string, row, column string) string {
	if table[row] != nil {
		if value, ok := table[row][column]; ok {
			return value
		}
	}
	return ""
}
