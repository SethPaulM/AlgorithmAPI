package main

import (
	"fmt"
	// "net/http"
)

func main() {
	// InsertionSort()
	MergeSort()
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
