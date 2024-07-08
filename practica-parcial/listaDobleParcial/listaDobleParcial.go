package main

type Nodo struct {

	valor int
	siguiente *Nodo

}

type ColaDoble struct {

	inicio *Nodo
	fin *Nodo

}	

// i)

func (c1 *ColaDoble) equals(c2 *ColaDoble) bool {   //retorna true si las dos colas son iguales

	nodo1 := c1.inicio
	nodo2 := c2.inicio
	
	
	if nodo1 == nil && nodo2 == nil {
		return false								//retorna false debido a que las dos colas estan vacias
	} else {

		for nodo1 != nil && nodo2 != nil {

			if nodo1.valor != nodo2.valor {
				return false							// si alguno de los valores difiere retorna false
			}

			nodo1 = nodo1.siguiente
			nodo2 = nodo2.siguiente

		}

	}

	return nodo1 == nil && nodo2 == nil     		// en este caso va a retornar true, si ambas listas terminaron.
													// si no terminaron retorna false, ya que una de las condiciones se cumple y la otra no.
}

func (c1 *ColaDoble) empty() bool {
	return c1.inicio == nil				// si c1.inicio == nil quiere decir que la cola no contiene elementos, por lo tanto esta premisa es verdadera
}										// en caso contrario, si hay algun elemento c1.inicio == nil es falso.

func (c1 ColaDoble) tamaño() int {

	if c1.inicio == nil {
		
		return 0				// la lista no contiene elementos

	} else {
		
		c1.inicio = c1.inicio.siguiente
		return 1 + c1.tamaño()

	}

}
