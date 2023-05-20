package src

// A productions with conditions expressed using features
type argument struct {
	Symbol string
	Production string
	DotPos int
}

// Returns a new argument type
func newArg(s string, p string, d int) argument {
	return argument {
		Symbol: s,
		Production: p,
		DotPos: d,
	}
}

// Function to get the symbol the dot is pointing to
func ( a argument ) GetTransition() string {
	if a.Production == "ε" || a.DotPos == -1 {
		return "ε"
	} else {
		return string(a.Production[a.DotPos])
	}
}

// Returns the same argument but changes the dot position
func ( a argument ) NextArgument() argument {
	arg := a
	if arg.DotPos == len(arg.Production)-1  || a.Production == "ε" {
		arg.DotPos = -1
		return arg
	}
	arg.DotPos++
	return arg
}

// Returns a map with a symbol and a single production each
func (g grammar) Argumented() map[argument]int {
	arguments := make(map[argument]int )
	arguments[newArg(g.S + "'", g.S, 0)] = 0
	index := 1
	for symbol, productions := range g.P {
		for _,  production := range productions {
			if production == "ε" {
				arguments[newArg(symbol, production, -1)] = index
			} else {
				arguments[newArg(symbol, production, 0)] = index
			}
			index++
		}
	}
	return arguments
}

// Returns the first agument of the grammr that is of the form S' -> S
func (g grammar) GetFirstArgument() argument {
	return argument { g.S+"'", g.S, 0 }
}
