package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func createGrammar() {
	rand.Seed(time.Now().UnixNano())

	letrasMayusculas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letrasMinusculas := "abcdefghijklmnopqrstuvwxyz"

	// Generar un número aleatorio entre 1 y 3 para determinar la cantidad de letras
	cantidad := rand.Intn(3) + 1

	// Crear un slice para almacenar las letras generadas
	letrasAleatorias := make([]byte, cantidad)

	for i := 0; i < cantidad; i++ {
		letra := generarLetraAleatoria(letrasMayusculas, letrasAleatorias[:i])
		letrasAleatorias[i] = letra
	}

	// Generar las producciones utilizando las letras generadas y letras minúsculas
	for _, letra := range letrasAleatorias {
		producciones := generarProducciones(letra, letrasAleatorias, letrasMinusculas)
		fmt.Printf("%s -> %s\n", string(letra), strings.Join(producciones, " | "))
	}
}

func generarLetraAleatoria(letrasMayusculas string, letrasGeneradas []byte) byte {
	for {
		letra := letrasMayusculas[rand.Intn(len(letrasMayusculas))]
		repetida := false

		// Verificar si la letra ya ha sido generada
		for _, letraGenerada := range letrasGeneradas {
			if letra == letraGenerada {
				repetida = true
				break
			}
		}

		// Si la letra no está repetida, retornarla
		if !repetida {
			return letra
		}
	}
}

func generarProducciones(letra byte, letrasGeneradas []byte, letrasMinusculas string) []string {
	var producciones []string

	// Generar un número aleatorio entre 0 y 1 para determinar si se incluye una producción que derive en ε
	incluirEpsilon := rand.Intn(2)

	// Generar un número aleatorio entre 1 y 3 para determinar la cantidad de producciones
	cantidadProducciones := rand.Intn(3) + 1

	// Generar producciones con 1 a 3 caracteres entre terminales y no terminales
	for i := 0; i < cantidadProducciones; i++ {
		var opciones []string

		// Generar un número aleatorio entre 1 y 3 para determinar la cantidad de caracteres en la producción
		cantidadCaracteres := rand.Intn(3) + 1

		for j := 0; j < cantidadCaracteres; j++ {
			if len(letrasGeneradas) > 0 && rand.Intn(2) == 0 {
				// Usar una letra generada aleatoria
				indice := rand.Intn(len(letrasGeneradas))
				opciones = append(opciones, string(letrasGeneradas[indice]))
			} else {
				// Usar una letra minúscula aleatoria
				indice := rand.Intn(len(letrasMinusculas))
				opciones = append(opciones, string(letrasMinusculas[indice]))
			}
		}

		produccion := strings.Join(opciones, "")
		producciones = append(producciones, produccion)
	}

	// Verificar si se debe incluir la producción que derive en ε
	if cantidadProducciones < 3 && cantidadProducciones >= 1 {
		if incluirEpsilon == 1 {
			producciones = append(producciones, "ε")
		}
	}

	return producciones
}