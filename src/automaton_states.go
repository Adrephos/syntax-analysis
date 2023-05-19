package src

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

// A state of a LR(0) automaton with kernel and clossure
type state struct {
	Kernel sets.Set[argument]
	Clossure sets.Set[argument]
}

// Returns a new state for an LR(0) automaton for a given kernel set
func NewState(g grammar, k sets.Set[argument]) state {
	kernel := k		
	clossure := k.Clone()	
	var nTerminal []string
	alreadyVisited := sets.NewString()
	for item := range clossure {
		transition := item.GetTransition()
		if g.N.Has(transition) {
			nTerminal = append(nTerminal, transition)
		}
	}
	if len(nTerminal) == 0 { return state{ kernel, clossure }}

	index := 0; finishFlag := true

	// Calculates clossure
	for finishFlag {
		current := nTerminal[index]
		if alreadyVisited.Has(current) {
			if index == len(nTerminal) - 1 {
				finishFlag = false
			} else {
				index++
			}
			continue
		} else {
			alreadyVisited.Insert(current)
		}

		for _, production := range g.P[current] {
			num := 0
			if production == "ε" {
				num = -1
			}
			arg := newArg(current, production, num)
			transition := arg.GetTransition()
			clossure.Insert(arg)

			if g.N.Has(transition) && !alreadyVisited.Has(transition) {
				nTerminal = append(nTerminal, transition)
			}
		}
		if index == len(nTerminal) - 1 {
			finishFlag = false
		} else {
			index++
		}
	}

	return state{ kernel, clossure }
}
