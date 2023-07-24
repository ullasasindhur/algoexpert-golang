package heapsort

import "fmt"

func Heap_Sort(input *[]int) {
	fmt.Println("Heap Sort")
	heapify(input, 0, len(*input)-1)
	fmt.Println(*input)
}

func heapify(input *[]int, start int, end int) {
	if end < 0 {
		return
	}
	var i, left, right, current int
	for i = end / 2; i >= 0; i-- {
		left = (2 * i) + 1
		right = (2 * i) + 2
		current = i
		if left <= end && (*input)[left] > (*input)[current] {
			current = left
		}
		if right <= end && (*input)[right] > (*input)[current] {
			current = right
		}
		if current != i {
			(*input)[current], (*input)[i] = (*input)[i], (*input)[current]
		}
	}
	(*input)[end], (*input)[start] = (*input)[start], (*input)[end]
	heapify(input, start, end-1)
}
