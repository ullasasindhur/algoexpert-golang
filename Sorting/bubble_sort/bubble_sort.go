package bubble_sort

import (
	"fmt"
)

func Bubble_sort(array *[]int) {
	fmt.Println("Bubble Sort: ")
	var arrayLen int = len(*array)
	var counter int
	var isSorted bool
	for !isSorted {
		isSorted = true
		for i := 0; i < arrayLen-1-counter; i++ {
			if (*array)[i] > (*array)[i+1] {
				(*array)[i], (*array)[i+1] = (*array)[i+1], (*array)[i]
				isSorted = false
			}
		}
		counter++
	}
	fmt.Println(*array)
}
