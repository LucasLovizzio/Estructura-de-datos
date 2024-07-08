// Description: Programa que realiza una cuenta regresiva desde un número dado hasta 0.

package main

import "fmt"

func cuenta_regresiva(num int) {
	if num < 0 {
		return
	}
	fmt.Println(num)
	cuenta_regresiva(num - 1)
}

func main() {
	fmt.Println("Cuenta regresiva: ")
	cuenta_regresiva(10)
}
