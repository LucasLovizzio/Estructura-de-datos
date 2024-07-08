/* Escriba un programa que lea una secuencia de caracteres e informe si la misma es
palíndromo. Utilice en la solución los TADś implementados. */

package main

import (
	"fmt"
)

// Pila

type Nodo struct {
	valor         string
	siguienteNodo *Nodo
}

type Pila struct {
	top  *Nodo
	size int
}

func (p *Pila) Push(valor string) {
	p.top = &Nodo{
		valor:         valor,
		siguienteNodo: p.top,
	}
	p.size++
}

func (p *Pila) Pop() string {
	valor := p.top.valor
	p.top = p.top.siguienteNodo
	p.size--
	return valor
}

func (p *Pila) IsEmpty() bool {
	return p.top == nil
}

func (p *Pila) GetItem() string {
	return p.top.valor
}

func (p *Pila) GetSize() int {
	return p.size
}

// Codigo

func main() {

	var palabra string

	var pila Pila

	fmt.Println("Ingrese una secuencia de caracteres: ")
	fmt.Scan(&palabra)

	for i := 0; i < len(palabra); i++ {
		pila.Push(string(palabra[i]))
	}
	palabraInvertida := ""

	for !pila.IsEmpty() {
		palabraInvertida += pila.Pop()
	}

	if palabra == palabraInvertida {
		fmt.Println("La palabra es palíndromo")
	} else {
		fmt.Println("La palabra no es palíndromo")
	}
}
