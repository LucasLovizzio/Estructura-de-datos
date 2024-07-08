/* Escriba un programa que lea 10 números y los imprima en orden inverso.
¿Utilizó en la solución alguno de los TAD’s que implementó en esta práctica? Si
tuviera que imprimir ahora los números en el orden que fueron leídos,
¿qué TAD de los que ya implementó usaría? */

package main

import (
	"fmt"
	"cola_pila"
)

func main() {
	var cola cola_pila.Cola
	var pila cola_pila.Pila
	fmt.Println("A continuación se le pedirá que ingrese 10 números.")
	for i := 0; i < 10; i++ {
		var n int
		fmt.Print("Ingrese un número: ")
		fmt.Scan(&n)
		pila.Push(n)
		cola.Encolar(n)
	}

	fmt.Println("Imprimiendo en orden inverso:")
	for !pila.IsEmpty() {
		fmt.Println(pila.Pop())
	}

	fmt.Println("Imprimiendo en orden de ingreso:")
	for !cola.IsEmpty() {
		fmt.Println(cola.Desencolar())
	}
}
