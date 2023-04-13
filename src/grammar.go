package src

import "k8s.io/apimachinery/pkg/util/sets"

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
