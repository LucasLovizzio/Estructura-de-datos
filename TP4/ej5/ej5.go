/* Escriba un programa que determine si una secuencia de llaves, corchetes y
paréntesis esta balanceada. Es decir, por cada símbolo de apertura debe aparecer
su correspondiente símbolo de cierre. Tenga en cuenta que los símbolos pueden
aparecer anidados. */

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


func main() {

	var secuencia string

	var pila Pila

	fmt.Println("Ingrese una secuencia de llaves, corchetes y paréntesis: ")
	fmt.Scan(&secuencia)

	for _, simbolo := range secuencia {
		if string(simbolo) == "{" || string(simbolo) == "[" || string(simbolo) == "(" {
			pila.Push(string(simbolo))
		} else if string(simbolo) == "}" {
			if pila.IsEmpty() || pila.GetItem() != "{" {
				fmt.Println("La secuencia no está balanceada.")
				return
			} else {
				pila.Pop()
			}
		} else if string(simbolo) == "]" {
			if pila.IsEmpty() || pila.GetItem() != "[" {
				fmt.Println("La secuencia no está balanceada.")
				return
			} else {
				pila.Pop()
			}
		} else if string(simbolo) == ")" {
			if pila.IsEmpty() || pila.GetItem() != "(" {
				fmt.Println("La secuencia no está balanceada.")
				return
			} else {
				pila.Pop()
			}
		}
	}

	if pila.IsEmpty() {
		fmt.Println("La secuencia está balanceada.")
	} else {
		fmt.Println("La secuencia no está balanceada.")
	}

}