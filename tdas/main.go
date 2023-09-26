package main

import (
	"fmt"
	TDALista "tdas/lista"
)

/* func main() {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	lista.Iterar(func(v int) bool {
		fmt.Println(v)

		return true

	})
} */

func main() {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)

	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {

		*contador_ptr += v
		return true
	})

	fmt.Println(*contador_ptr)

	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		fmt.Println(iter.VerActual())
	}
	var (
		suma    int64                 = 0
		puntero *int64                = &suma
		lista2  TDALista.Lista[int64] = TDALista.CrearListaEnlazada[int64]()
	)
	lista2.InsertarPrimero(int64(2))
	lista2.InsertarPrimero(int64(6))
	lista2.InsertarPrimero(int64(8))
	//iteramos haciendo una suma
	for iter := lista2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		*puntero += iter.VerActual()
	}
	fmt.Println(suma)

}
