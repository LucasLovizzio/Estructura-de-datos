package main

import "fmt"

func posicion_mayus(texto string, indice int) int {
	if indice >= len(texto) {
		return -1
	}

	if texto[indice] >= 'A' && texto[indice] <= 'Z' {
		return indice
	}

	return posicion_mayus(texto, indice+1)
}

func main() {
	texto := "hola Mundo"
	fmt.Println(posicion_mayus(texto, 0))
}
