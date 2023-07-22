package insertion_sort

import (
	"fmt"
)

func Insertion_sort(array *[]int) {
	fmt.Println("Insertion Sort: ")
	var arrayLen int = len(*array)
	var i, j int
	for i = 1; i < arrayLen; i++ {
		j = i
		for j > 0 && (*array)[j-1] > (*array)[j] {
			(*array)[j-1], (*array)[j] = (*array)[j], (*array)[j-1]
			j--
		}
	}
	fmt.Println(*array)
}
