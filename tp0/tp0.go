package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	var aux int = *x
	*x = *y
	*y = aux
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) != 0 {
		var aux [2]int = [2]int{vector[0], 0}
		for i := 1; i < len(vector); i++ {
			if vector[i] > aux[0] {
				aux[0] = vector[i]
				aux[1] = i
			}
		}
		return aux[1]
	}
	return -1
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	var i int
	if len(vector1) > len(vector2) {
		if len(vector2) == 0 {
			return 1
		}
		for i = 0; i < len(vector2); i++ {
			if vector1[i] > vector2[i] {
				return 1
			} else if vector1[i] < vector2[i] {
				return -1
			}
		}
		return 1
	} else if len(vector1) < len(vector2) {
		if len(vector1) == 0 {
			return -1
		}
		for i = 0; i < len(vector1); i++ {
			if vector1[i] > vector2[i] {
				return 1
			} else if vector1[i] < vector2[i] {
				return -1
			}
		}
		return -1
	} else if len(vector1) == len(vector2) {
		for j := 0; j < len(vector1); j++ {
			if vector1[j] > vector2[j] {
				return 1
			} else if vector1[j] < vector2[j] {
				return -1
			}
		}
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	var max int
	if len(vector) > 1 {
		for i := 1; i <= len(vector)-1; i++ {
			max = Maximo(vector[0 : len(vector)-i])

			if vector[max] > vector[len(vector)-i] {
				Swap(&vector[max], &vector[len(vector)-i])
			}
		}
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 1 {
		return vector[0]
	} else if len(vector) == 0 {
		return 0
	}
	return Suma(vector[:len(vector)-1]) + vector[len(vector)-1]
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	largo := len(cadena)
	if 1 < largo && largo <= 3 {
		if cadena[0] == cadena[largo-1] {
			return true
		} else {
			return false
		}
	} else if largo > 3 {
		if EsCadenaCapicua(cadena[1:largo-1]) && cadena[0] == cadena[largo-1] {
			return true
		} else {
			return false
		}
	}
	return true
}
