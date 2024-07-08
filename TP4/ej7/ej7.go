/* . Implemente una cola doblemente terminada usando nodos. Las operaciones
encolar() y desencolar() deben tener O(1). */

package main

type Node struct {
	valor int
	siguiente *Node
	anterior *Node
}

type ListaDoblementeEnlazada struct {
	cabeza *Node
	cola *Node
}

func (list * ListaDoblementeEnlazada) encolar(valor int) {
	
	nuevo := &Node{valor: valor, siguiente: nil, anterior: nil}

	if list.cabeza == nil {
		list.cabeza = nuevo
		list.cola = nuevo
	} else {
		list.cola.siguiente = nuevo
		nuevo.anterior = list.cola
		list.cola = nuevo
	}
}

func (list *ListaDoblementeEnlazada) desencolar() int {
	
	if list.cabeza == nil {
		return -1
	}

	valor := list.cabeza.valor
	list.cabeza = list.cabeza.siguiente
	if list.cabeza != nil {
		list.cabeza.anterior = nil
	} else {
		list.cola = nil
	}
	return valor
}
