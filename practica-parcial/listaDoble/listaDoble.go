package main

type Nodo struct {
	valor      	int
	anterior 	*Nodo
	siguiente 	*Nodo
}

type ListaDoble struct {
	cabeza 	*Nodo
	cola 	*Nodo
}

func (lista *ListaDoble) encolar(num int) {

	nuevoNodo := &Nodo{valor : num, anterior : nil, siguiente : nil}		// Creamos el nuevo nodo, con el valor dado por parametro

	if lista.cabeza == nil {			// En una lista, se recorren los elementos desde la cabeza y se ingresan en el primer lugar disponible
										// que generalmente es el siguiente del ultimo. Por lo tanto, si la cabeza esta vacia, quiere decir que no hay
		lista.cabeza = nuevoNodo		// elementos creados y por lo tanto el primero va a ser el que ingresemos.
		lista.cola = nuevoNodo			// Nuestro elemento es el primero y el ultimo.

	} else {									// si ya tenemos elementos pasa lo siguiente : el nodo se va a poner en la cola (ultimo lugar)
												// por lo tanto el nodo anterior al nuevonodo, va a ser el que actualmente es el ultimo.
		nuevoNodo.anterior = lista.cola			// el siguiente al ultimo nodo actual va a ser nuestro nuevo nodo. Entonces, ahora si, el ultimo elemento
		lista.cola.siguiente = nuevoNodo		// (cola) va a ser nuestro nuevo nodo.
		lista.cola = nuevoNodo

	}
}

func (lista *ListaDoble) agregar_frente( num int ) {

	nuevoNodo := &Nodo{valor : num, anterior: nil, siguiente: nil}

	if lista.cabeza == nil {

		lista.cabeza = nuevoNodo
		lista.cola = nuevoNodo

	} else {

		nuevoNodo.siguiente = lista.cabeza
		lista.cabeza.anterior = nuevoNodo
		lista.cabeza = nuevoNodo

	}

}

func (lista1 *ListaDoble) igualdad(lista2 *ListaDoble) bool {

	nodo1 := lista1.cabeza				// tomamos los primeros elementos de cada lista
	nodo2 := lista2.cabeza

	for nodo1 != nil && nodo2 != nil {		

		if nodo1.valor != nodo2.valor {

			return false

		}

		nodo1 = nodo1.siguiente
		nodo2 = nodo2.siguiente

	}

	return nodo1 == nil && nodo2 == nil 		// Si ambos nodos son nil al mismo tiempo, las listas son iguales

}

func (lista ListaDoble) tamaño() int {

	if lista.cabeza == nil {				// si no hay un primer elemento en la lista, la lista esta vacia.

		return 0								// tamaño = 0

	} else {												// si hay un primer elemento en la lista

		lista.cabeza = lista.cabeza.siguiente  				// contamos ese y lo eliminamos, el primer elemento ahora va a ser el segundo
		return 1 + lista.tamaño()							// vamos a ir contando elementos y sacando hasta que nos quedemos sin elementos
															// por lo tanto volvera a if (lista.cabeza == nil), lo que retornara 0 y devolvera la suma.
	}

}

func (lista ListaDoble) tamaño2() int {

	return tamañoR( lista.cabeza ) 

}

func tamañoR ( nodo * Nodo ) int {

	if nodo == nil {

		return 0
		
	} else {

		return 1 + tamañoR(nodo.siguiente)

	}

}


