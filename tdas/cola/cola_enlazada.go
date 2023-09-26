package cola

const cola_vacia string = "La cola esta vacia"

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	principio *nodo[T]
	fin       *nodo[T]
}

func crearNodo[T any](elemento T, sig *nodo[T]) *nodo[T] {
	return &nodo[T]{dato: elemento, siguiente: sig}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	if c.principio == nil {
		return true
	}
	return false
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(cola_vacia)
	}
	return c.principio.dato

}

func (c *colaEnlazada[T]) Encolar(elem T) {
	nuevo_nodo := crearNodo(elem, nil)
	if c.EstaVacia() {
		c.principio = nuevo_nodo
		c.fin = nuevo_nodo
		return
	}
	c.fin.siguiente = nuevo_nodo
	c.fin = nuevo_nodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	var dato_a_devolver T
	if c.EstaVacia() {
		panic(cola_vacia)
	}

	dato_a_devolver = c.VerPrimero()

	if c.principio != c.fin {
		c.principio = c.principio.siguiente
	} else {
		c.principio = nil
		c.fin = nil
	}
	return dato_a_devolver
}
