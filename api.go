package main

import (
	"fmt"
	// "net/http"
	// "strconv"
)

func main() {
	// InsertionSort()
	// MergeSort()
	QuickSort([8]int{2, 8, 7, 1, 3, 5, 6, 4}, 0 , 7)
}

func InsertionSort() {
	var array = [6]int{5, 2, 4, 6, 1, 3}
	for i := 0; i < len(array); i++ {
		var currentElement = array[i]
		var lastIndex = i - 1
	  
		for lastIndex >= 0 && array[lastIndex] > currentElement {
			array[lastIndex + 1] = array[lastIndex]
			lastIndex--
		}
		array[lastIndex + 1] = currentElement
	}
    fmt.Println(array)
}

func MergeSort() {
	var array = [8]int{2, 4, 6, 7, 1, 2, 3, 5}
	var left_array = [4]int{}
	var right_array = [4]int{}

	for i := 0; i < 4; i++ {
		left_array[i] = array[i]
	}
	for j := 4; j < len(array); j++ {
		right_array[(j - 4)] = array[j]	
	}

	var i int = 0 // iterate for left_array
	var j int = 0 // iterate for right_array
	var k int = 0 // iterate for array

	for i < len(left_array) && j < len(right_array) {
		if left_array[i] <= right_array[j] {
			array[k] = left_array[i]
			i = i + 1
		} else  {
			array[k] = right_array[j]
			j = j + 1
		}
		k = k + 1
	}

	for i < len(left_array) {
		array[k] = left_array[i]
		i = i + 1
		k = k + 1
	}
	
	for j < len(right_array) {
		array[k] = right_array[j]
		j = j + 1
		k = k + 1
	}

	fmt.Println(left_array)
	fmt.Println(right_array)
	fmt.Println(array)
}

func QuickSort(A [8]int, p int, r int) {
	if p < r {
		var x int = A[r]
		var i int = p - 1
		for j := p; j < r - 1; j++ {
			if A[j] <= x {
				i = i + 1
				A[i], A[j] = A[j], A[i]
			}
		}
		A[i + 1], A[r] = A[r], A[i + 1]
		i = i + 1
		QuickSort(A, p, i - 1)
		QuickSort(A, i + 1, r)
		fmt.Println(A)
	}
	
}

func Partition(A [8]int, p int, r int) int {
	var x int = A[r]
	var i int = p - 1
	for j := p; j < r - 1; j++ {
		if A[j] <= x {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i + 1], A[r] = A[r], A[i + 1]
	return i + 1
}

/* 
.A; p; r/
if p < r
2 // Partition the subarray around the pivot, which ends up in AŒq�.
3 q D PARTITION.A; p; r/
4 QUICKSORT.A; p; q  1/ // recursively sort the low side
5 QUICKSORT.A; q C 1; r/ // recursively sort the high side 

*/