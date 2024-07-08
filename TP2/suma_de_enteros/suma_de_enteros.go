package main

import "fmt"

func suma_de_enteros(num1 int, num2 int) int {
	if num1 == num2 {
		return num1
	}
	if num1 < num2 {
		return num1 + suma_de_enteros(num1+1, num2)
	}else{
		return num2 + suma_de_enteros(num1, num2+1)
	}
}

func main() {
	fmt.Println(suma_de_enteros(5, 12))
}
