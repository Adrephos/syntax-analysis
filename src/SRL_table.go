package src

import (
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
)

// SLR Table with states and grammar symbols
type table map[int]map[string]fn

// Type to define function used in SLR parsing table
type fn struct {
	function string // Whether is  Goto(g), shift(s) or reduce (r)
	number int // The state or grammar production the function refers to
}

// add a new value at table[row][column]
func ( t table) addCell(row int, column string, value fn) error {
	if t.getValue(row, column) != (fn{}) {
		errorMsg := fmt.Sprintf("Error: Conflicto en la celda (%v, %v) ya existe un valor", row, column)
		return errors.New(errorMsg)
	}
	if t[row] == nil {
		t[row] = make(map[string]fn)
	}
	t[row][column] = value
	return nil
}

// gets a value from table[row][column]
func ( t table ) getValue(row int, column string) fn {
	if t[row] != nil {
		if value, ok := t[row][column]; ok {
			return value
		}
	}
	return fn{}
}

// If exists a state s in statesMap returns true and the number
func checkStateNumber(k sets.Set[argument], statesMap map[int]sets.Set[argument]) (int, bool){
	for key, value := range statesMap {
		if value.Difference(k).Len() == 0 {
			return key, true
		}
	}
	return -1, false
}

// From an automaton state makes all its transitions
func makeTansitions(t *table, s state, sM *map[int]sets.Set[argument], nextStateNumber *int, g grammar, arguments map[argument]int, followSets map[string]sets.String) error {
	var err error
	currentStateNumber,_ := checkStateNumber(s.Kernel, *sM)
	transitionsMap := make(map[string]sets.Set[argument])
	for item := range s.Clossure {
		if item.GetTransition() == "ε" {
			if item.Symbol == g.S+"'" {
				function := fn{"a", 0}
				err = (*t).addCell(currentStateNumber, "$", function)
				if err != nil {
					return err
				}
			} 
			dot := 0
			if item.Production == "ε" { dot = -1 }
			arg := newArg(item.Symbol, item.Production, dot)
			number := arguments[arg]
			for symbol := range followSets[item.Symbol] {
				function := fn{"r", number}
				err = (*t).addCell(currentStateNumber, symbol, function)
				if err != nil {
					return err
				}
			}
			continue
		}
		if _, ok := transitionsMap[item.GetTransition()]; ok {
			transitionsMap[item.GetTransition()] = transitionsMap[item.GetTransition()].Insert(item.NextArgument())
		} else {
			transitionsMap[item.GetTransition()] = sets.New(item.NextArgument())
		}
	}

	for symbol, kernel := range transitionsMap {
		number, exists := checkStateNumber(kernel, *sM)
		if !exists {
			number = *nextStateNumber
			(*sM)[number] = kernel
			(*nextStateNumber)++
		}
		if g.N.Has(symbol) {
			function := fn{"g", number}
			err = (*t).addCell(currentStateNumber, symbol, function)
			if err != nil {
				return err
			}
		} else if g.T.Has(symbol) {
			function := fn{"s", number}
			err = (*t).addCell(currentStateNumber, symbol, function)
			if err != nil {
				return err
			}
		} 
	}
	return nil
}

// Function to create LR(0) automaton and and besed on it return the SLR table
func (g grammar ) CreateSLRTable() (table, map[argument]int, error) {
	stateNumber := 1
	finish := false
	table := make(table)
	visited := sets.NewInt()
	arguments := g.Argumented()
	followSets := g.FollowGrammar()
	firstKernel := sets.New(g.GetFirstArgument())
	statesMap := make(map[int]sets.Set[argument])
	statesMap[0] = firstKernel

	for !finish {
		counter := 0
		for num, state := range statesMap {
			if visited.Has(num){
				counter++
				if stateNumber == counter  {
					finish = true
					break
				}
				continue
			}
			counter = 0
			err := makeTansitions(&table, NewState(g, state), &statesMap, &stateNumber, g, arguments, followSets)
			if err != nil {
				return nil, nil, err
			}
			visited.Insert(num)
		}
	}
	return table, arguments, nil
}
