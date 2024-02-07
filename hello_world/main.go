package main

import "fmt"

/* func main() {
	fmt.Println("Hello world")
}
*/

///MERGE SORT///
func mergeSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	mid := length / 2
	left := make([]int, mid)
	right := make([]int, length-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	mergeSort(left)
	mergeSort(right)

	merge(arr, left, right)
}

func merge(arr []int, left []int, right []int) {
	i, j, k := 0, 0, 0
	leftLength, rightLength := len(left), len(right)

	for i < leftLength && j < rightLength {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < leftLength {
		arr[k] = left[i]
		i++
		k++
	}

	for j < rightLength {
		arr[k] = right[j]
		j++
		k++
	}
}

///////////////

//QUICK SORT//
func QuickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := Partition(arr, low, high)
		QuickSort(arr, low, pivotIndex)
		QuickSort(arr, pivotIndex+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[low]
	left := low
	right := high

	for left < right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	return right
}

////////////////////////

/////SELECTION SORT/////

func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

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

////////////////////////

func mezclar(arr []int) {
	if len(arr) <= 2 {
		return
	}

	m := len(arr) / 2
	mezclar(arr[:m])
	mezclar(arr[m:])

	for i := 0; i < m; i++ {
		if i%2 == 1 {
			arr[i], arr[m+i-1] = arr[m+i-1], arr[i]
		}
	}
}

func countInRange(nums []int, num, left, right int) int {
	count := 0
	for i := left; i < right; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}

func majorityElementRecursive(nums []int, left, right int) int {
	// Caso base: si el rango tiene un solo elemento, ese es el elemento mayoritario
	if left == right {
		return nums[left]
	}

	// Dividir el rango en dos mitades
	mid := (left + right) / 2

	// Llamar recursivamente a las mitades izquierda y derecha
	leftMajority := majorityElementRecursive(nums, left, mid)
	rightMajority := majorityElementRecursive(nums, mid+1, right)

	// Contar la frecuencia de los elementos mayoritarios en cada mitad
	leftCount := countInRange(nums, leftMajority, left, right)
	rightCount := countInRange(nums, rightMajority, left, right)

	// Determinar el elemento mayoritario en todo el rango
	if leftCount > (right-left)/2 {
		return leftMajority
	} else if rightCount > (right-left)/2 {
		return rightMajority
	} else {
		return -1 // No hay un elemento mayoritario en este rango
	}
}

func majorityElement(nums []int) bool {
	// Llamar a la función principal con el rango completo del array
	result := majorityElementRecursive(nums, 0, len(nums)-1)

	// Verificar si el elemento obtenido es el mayoritario en todo el array
	return result != -1
}

func main() {
	//vec := []int{1, 3, 5, 7, 2, 4, 6, 8}
	//vec := []int{1, 3, 5, 7, 9, 11, 13, 15, 2, 4, 6, 8, 10, 12, 14, 16}
	//mezclar(vec)
	//fmt.Println(vec)

	vec := []int{1, 3, 4, 7, 1, 1, 6, 8, 1, 1, 1, 1, 1}
	fmt.Println(majorityElement(vec))
}
