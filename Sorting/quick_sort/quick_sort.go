package quicksort

import "fmt"

func Quick_sort(input *[]int) {
	fmt.Println("Quick Sort: ")
	quick(input, 0, len(*input)-1)
	fmt.Println(*input)
}

func quick(input *[]int, start int, end int) {
	if start >= end {
		return
	}
	var pivot int = (*input)[start]
	var i, j int = start, end
	for start <= end {
		if (*input)[start] > pivot && (*input)[end] < pivot {
			(*input)[start], (*input)[end] = (*input)[end], (*input)[start]
		}
		if (*input)[start] <= pivot {
			start++
		}
		if (*input)[end] >= pivot {
			end--
		}
	}
	(*input)[i], (*input)[end] = (*input)[end], (*input)[i]
	quick(input, i, end-1)
	quick(input, end+1, j)
}
