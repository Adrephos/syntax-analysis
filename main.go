package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/Adrephos/syntax-analysis/src"
)

func main() {

	argsWithoutProg := os.Args[1:]


	if len(argsWithoutProg) == 0 {

		grammar := src.GrammarInput()
		strings := src.StringsInput()

		grammar.Print()

		src.Ll1Routine(grammar, strings)
		src.Lr0Routine(grammar, strings)

	} else if argsWithoutProg[0] == "-g" {

		numCadenas, err := strconv.Atoi(argsWithoutProg[1])

		if err != nil {
			fmt.Println("Introduzca un número")
		}	else {
			src.CreateGrammar(numCadenas)
		}

	} else if argsWithoutProg[0] == "-f"{
		
		if len(argsWithoutProg) >= 2 {
			src.FileInput(argsWithoutProg[1])
		} else {
			fmt.Println("Ningún archivo fue proporcionado")
		}

	} else if argsWithoutProg[0] == "-h" {

		f, err := ioutil.ReadFile("./utils/help.txt")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(f))

	} else {
		fmt.Println("unknown option:", argsWithoutProg[0])
	}
}
