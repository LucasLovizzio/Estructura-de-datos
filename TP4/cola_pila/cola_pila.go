package cola_pila

/* Escriba un programa que lea 10 números y los imprima en orden inverso.
¿Utilizó en la solución alguno de los TAD’s que implementó en esta práctica? Si
tuviera que imprimir ahora los números en el orden que fueron leídos,
¿qué TAD de los que ya implementó usaría? */

type Nodo struct {
	valor         int
	siguienteNodo *Nodo
}

type Pila struct {
	top  *Nodo
	size int
}

func (p *Pila) Push(valor int) {
	p.top = &Nodo{
		valor:         valor,
		siguienteNodo: p.top,
	}
	p.size++
}

func (p *Pila) Pop() int {
	valor := p.top.valor
	p.top = p.top.siguienteNodo
	p.size--
	return valor
}

func (p *Pila) IsEmpty() bool {
	return p.top == nil
}

func (p *Pila) GetItem() int {
	return p.top.valor
}

func (p *Pila) GetSize() int {
	return p.size
}

type Cola struct{
	salida * Nodo;
	entrada * Nodo;
	size int;
}

func (c *Cola) Encolar(valor int) {
	nuevo := &Nodo{
		valor: valor,
		siguienteNodo:  nil,
	}

	if c.entrada == nil {
		c.entrada = nuevo
		c.salida = nuevo
	} else {
		c.entrada.siguienteNodo = nuevo
		c.entrada = nuevo
	}
}

func (c *Cola) Desencolar() int {
	if c.salida == nil {
		return 0
	}
	valor := c.salida.valor
	c.salida = c.salida.siguienteNodo
	return valor
}

func (c *Cola) IsEmpty() bool {
	return c.salida == nil
}

func (c *Cola) GetItem() int {
	return c.salida.valor
}

func (c *Cola) GetSize() int {
	return c.size
}
