// conjunto

package main

import ("fmt")

type Nodo struct {

	valor int
	siguiente * Nodo

}

type Conjunto struct {

	inicio * Nodo
	fin * Nodo

}

func (c *Conjunto) Contiene(valor_buscado int) bool {
    
	aux := c.inicio 

	for aux != nil {

		if aux.valor == valor_buscado {

			return true								// retorna true si el valor ya esta en el conjunto

		}

		aux = aux.siguiente	

	}

	return false									// retorna false si el valor no esta en el conjunto

}

func (c *Conjunto) encolar(valor int) {

	if c.Contiene(valor){

		return		// si el valor ya esta en la lista, no lo va a agregar ya que los conjuntos no repiten elementos.

	}

	nuevo_nodo := &Nodo{valor: valor, siguiente: nil}

	if c.inicio == nil {

		c.inicio = nuevo_nodo
		c.fin = nuevo_nodo

	} else {

		c.fin.siguiente = nuevo_nodo
		c.fin = nuevo_nodo

	}

}

func (c Conjunto) tamaño() int {

	if c.inicio == nil {

		return 0

	} else {

		c.inicio = c.inicio.siguiente
		return 1 + c.tamaño()

	}

}

func main() {

	nuevoConjunto := &Conjunto{inicio : nil}

	fmt.Println(nuevoConjunto.tamaño())

}