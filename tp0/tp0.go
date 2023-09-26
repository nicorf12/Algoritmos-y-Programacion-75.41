package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
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
/* func Comparar(vector1 []int, vector2 []int) int {
	if len(vector1) != 0 || len(vector2) != 0 {
		if len(vector1) == 0 && len(vector2) > 0 {
			return -1
		} else if len(vector1) > 0 && len(vector2) == 0 {
			return 1
		}

		var max int
		if len(vector1) >= len(vector2) {
			max = len(vector1)
		} else if len(vector1) < len(vector2) {
			max = len(vector2)
		}
		for j := 0; j < max-1; j++ {
			if vector1[j] > vector2[j] {
				return 1
			} else if vector1[j] < vector2[j] {
				return -1
			}
		}
		if len(vector1) == max && !(len(vector1) == len(vector2)) {
			return 1
		} else if len(vector2) == max && !(len(vector1) == len(vector2)) {
			return -1
		}
	}
	return 0
} */

func Comparar(vector1 []int, vector2 []int) int {
	var min int
	if len(vector1) >= len(vector2) {
		min = len(vector2)
	} else {
		min = len(vector1)
	}

	for j := 0; j < min; j++ {
		if vector1[j] > vector2[j] {
			return 1
		} else if vector1[j] < vector2[j] {
			return -1
		}
	}

	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
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
	if largo <= 1 || (EsCadenaCapicua(cadena[1:largo-1]) && cadena[0] == cadena[largo-1]) { //quizas aqui este sacrificando claridad por menos lineas de codigo en esta
		return true
	} else {
		return false
	}
}
