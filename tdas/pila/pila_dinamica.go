package pila

const (
	parametro_de_creada = 1
	nro_de_redimension  = 2
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, parametro_de_creada), cantidad: 0}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	if p.cantidad == 0 {
		return true
	} else {
		return false
	}
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() == false {
		return p.datos[p.cantidad-1]
	} else {
		panic("La pila esta vacia")
	}
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(nro_de_redimension * cap(p.datos))
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if (2*nro_de_redimension)*p.cantidad <= cap(p.datos) {
		p.redimensionar(cap(p.datos) / nro_de_redimension)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

func (p *pilaDinamica[T]) redimensionar(capacidad_nueva int) {
	aux := make([]T, capacidad_nueva)
	copy(aux, p.datos)
	p.datos = aux
}
