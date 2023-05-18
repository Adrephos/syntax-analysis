package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func stringToGrammar(s string) grammar {
	productions := make(map[string][]string)
	initial := ""
	s = strings.ReplaceAll(s, " ", "")
	s = strings.TrimSpace(s)
	prodArr := strings.Split(s, "\n")
	
	for index, line := range prodArr {
		sy_prod := strings.Split(line, "->")
		if index == 0 { initial = sy_prod[0] }
		production := strings.Split(sy_prod[1], "|")
		productions[sy_prod[0]] = append(productions[sy_prod[0]], production...)

	}
	return mapToGrammar(productions, initial)
}

func FileInput(path string) {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	file := string(f)

	grammars_n_strings := strings.Split(file, "---\n")

	for i, input := range grammars_n_strings {
		inputArr := strings.Split(input, "\n\n")
		noSpaceStrings := strings.TrimSpace(inputArr[1])

		grammar := stringToGrammar(inputArr[0])
		stringsArr := strings.Split(noSpaceStrings, "\n")
		
		grammar.Print()

		Ll1Routine(grammar, stringsArr)
		Lr0Routine(grammar, stringsArr)
		
		if i != len(grammars_n_strings)-1 {
			fmt.Print("\n====================================================\n\n")
		}
	}
}
