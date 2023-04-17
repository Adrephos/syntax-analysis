package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"k8s.io/apimachinery/pkg/util/sets"
)

func GrammarInput() grammar {
	productions := make(map[string][]string)
	var line, initial string
	fmt.Println("Formato: <no-terminal> -> <production> | <production> .....")
	fmt.Println("Escribe $ para finalizar")
	count := 0
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line = scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		if line == "$" {
			break
		}
		sy_prod := strings.Split(line, "->")
		production := strings.Split(sy_prod[1], "|")
		productions[sy_prod[0]] = production
		if count == 0 {
			initial = sy_prod[0]
		}
		count++
	}

	return mapToGrammar(productions, initial)
}

func mapToGrammar( m map[string][]string, initial string) grammar {
	var g grammar
	g.P = m
	g.S = initial
	t,n := sets.NewString(), sets.NewString()

	for key := range m {
		n.Insert(key)
	}

	for _, value := range m {
		for _, production := range value {
			for _, char := range production {
				if !n.Has(string(char)) && string(char) != "ε" {
					t.Insert(string(char))
				}
			}
		}
	}
	g.T = t
	g.N = n

	return g
}
