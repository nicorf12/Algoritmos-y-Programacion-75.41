package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
///MERGE SORT///
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := make([]int, middle)
	right := make([]int, len(arr)-middle)

	copy(left, arr[:middle])
	copy(right, arr[middle:])

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
///////////////

//QUICK SORT//
func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
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
