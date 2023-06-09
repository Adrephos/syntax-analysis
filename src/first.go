package src

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/util/sets"
)

// Returns the first set for a Non terminal of a given grammar
func FirstNonTerminal(g grammar, s string, r sets.String) sets.String {
	first := sets.NewString()
	if g.T.Has(s) || s == "ε" {
		first.Insert(s)
		return first
	}
	// If s in non-terminal
	for _, production := range g.P[s] {
		firstSymbol := string(production[0])
		if production == "ε" {
			first.Insert(production)
		} else if g.T.Has(firstSymbol) {
			first.Insert(firstSymbol)
		}
		if g.N.Has(firstSymbol) && firstSymbol != s {
			epsilon := true
			for _, symbol := range production {
				symbolStr := string(symbol)

				if r.Has(symbolStr) { continue }

				symbolStrFirst := FirstNonTerminal(g, symbolStr, r.Insert(symbolStr))

				if !epsilon { break }

				if !symbolStrFirst.Has("ε") { epsilon = false }
				first = first.Union(symbolStrFirst.Delete("ε"))
			}
			if epsilon {
				first.Insert("ε")
			}
		}
	}

	return first
}

// The first set of a grammar symbol
func First(g grammar, s string) sets.String {
	first := sets.NewString()
	//If s is terminal or ε First(s) = {s}
	if g.T.Has(s) || s == "ε" {
		first.Insert(s)
		return first
	}
	if g.N.Has(s) {
		return FirstNonTerminal(g, s, sets.NewString(s))
	} else {
		log.Fatalf("Symbol '%v' is not part of this grammar.", s)
		return sets.NewString()
	}
}

// The first set of a given string
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

// Returns the first sets for each terminal and non terminal of a grammar
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

// Prints the first sets of a gramar for each symbol
func PrintFirstGrammar(fG map[string]sets.String) {
	for symbol, first := range fG {
		fmt.Println(symbol, "->", first.List())
	}
} 
