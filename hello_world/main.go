package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
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
