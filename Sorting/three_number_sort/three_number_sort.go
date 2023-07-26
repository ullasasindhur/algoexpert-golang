package threenumbersort

import "fmt"

func Three_Number_Sort(input *[]int, order *[]int) {
	fmt.Println("Three Number Sort:")
	var firstElement, secondElement int = (*order)[0], (*order)[1]
	var currentIdx, secondIdx, firstIdx int = 0, len(*input) - 1, 0
	var value int
	for firstIdx <= secondIdx {
		value = (*input)[firstIdx]
		if value == firstElement {
			(*input)[firstIdx], (*input)[currentIdx] = (*input)[currentIdx], (*input)[firstIdx]
			currentIdx++
			firstIdx++
		} else if value == secondElement {
			firstIdx++
		} else {
			(*input)[secondIdx], (*input)[firstIdx] = (*input)[firstIdx], (*input)[secondIdx]
			secondIdx--
		}
	}
	fmt.Println(*input)
}
