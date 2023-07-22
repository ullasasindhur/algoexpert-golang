package selection_sort

import (
	"fmt"
)

func Selection_sort(input *[]int) {
	fmt.Println("Selection Sort: ")
	var inputLen int = len(*input)
	var i, j, index int
	for i = 1; i < inputLen; i++ {
		index = i - 1
		for j = i; j < inputLen; j++ {
			if (*input)[index] > (*input)[j] {
				index = j
			}
		}
		if index != i-1 {
			(*input)[index], (*input)[i-1] = (*input)[i-1], (*input)[index]
		}
	}
	fmt.Println(*input)
}
