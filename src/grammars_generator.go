package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CreateGrammar(numCadenas int) {
	rand.Seed(time.Now().UnixNano())

	letrasMayusculas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letrasMinusculas := "abcdefghijklmnopqrstuvwxyz"

	// Crear un mapa para almacenar las producciones
	gramatica := make(map[byte][]string)

	// Generar un número aleatorio entre 1 y 3 para determinar la cantidad de letras
	cantidad := rand.Intn(3) + 1

	// Crear un slice para almacenar las letras generadas
	letrasAleatorias := make([]byte, cantidad)

	for i := 0; i < cantidad; i++ {
		letra := generarLetraAleatoria(letrasMayusculas, letrasAleatorias[:i])
		letrasAleatorias[i] = letra
	}

	var inicial byte 
	// Generar las producciones utilizando las letras generadas y letras minúsculas
	for i, letra := range letrasAleatorias {
		if i == 0 { inicial = letra }
		producciones := generarProducciones(letra, letrasAleatorias, letrasMinusculas)
		gramatica[letra] = producciones
		fmt.Printf("%s -> %s\n", string(letra), strings.Join(producciones, " | "))
	}

	// Generar varias cadenas aleatorias no repetidas y de máximo 10 caracteres
	cadenasGeneradas := make(map[string]bool)

	fmt.Println()
	count := 0
	for len(cadenasGeneradas) < numCadenas {
		cadena := generarCadenaAleatoria(inicial, gramatica, 0)
		if cadenasGeneradas[cadena] {
			count++
		}
		if len(cadena) <= 10 && !cadenasGeneradas[cadena] {
			cadenasGeneradas[cadena] = true
			fmt.Println(cadena)
		} 
		if count >= 1000 {
			fmt.Println("No es posible generar más cadenas de longitud 10")
			break
		}
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

func generarCadenaAleatoria(simbolo byte, gramatica map[byte][]string, i int) string {
	producciones, existe := gramatica[simbolo]
	if !existe {
		return ""
	}

	produccion := producciones[rand.Intn(len(producciones))]
	var cadena strings.Builder

	for _, caracter := range produccion {
		if caracter >= 'A' && caracter <= 'Z' {
			if i <= 20 {
				cadena.WriteString(generarCadenaAleatoria(byte(caracter), gramatica, i+1))
			} 
		} else if caracter != 'ε' && !(caracter >= 'A' && caracter <= 'Z') {
			cadena.WriteByte(byte(caracter))
		}
	}

	return cadena.String()
}
