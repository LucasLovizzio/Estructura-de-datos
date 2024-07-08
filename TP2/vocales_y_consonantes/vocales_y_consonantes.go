package main

import "fmt"

func vocales_y_consonantes(texto string) (int, int) {
	if texto == "" {
		return 0, 0
	}
	switch texto[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		vocales, consonantes := vocales_y_consonantes(texto[1:])
		return vocales + 1, consonantes
	case ' ':
		vocales, consonantes := vocales_y_consonantes(texto[1:])
		return vocales, consonantes
	default:
		vocales, consonantes := vocales_y_consonantes(texto[1:])
		return vocales, consonantes + 1
	}
}

func main() {
	vocales, consonantes := vocales_y_consonantes("hola mundo")
	fmt.Println("Vocales:", vocales)
	fmt.Println("Consonantes:", consonantes)
}
