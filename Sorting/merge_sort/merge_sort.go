package merge_sort

import "fmt"

func Merge_sort(input *[]int) {
	fmt.Println("Merge Sort:")
	split(input, 0, len(*input)-1)
	fmt.Println(*input)
}

func split(input *[]int, start int, end int) {
	if start != end {
		var mid int = (start + end) / 2
		split(input, start, mid)
		split(input, mid+1, end)
		merge(input, start, mid, end)
	}
}

func merge(input *[]int, start int, mid int, end int) {
	var arr1Len, arr2Len int = mid - start + 1, end - mid
	var arr1, arr2 []int = make([]int, arr1Len), make([]int, arr2Len)
	var i, j int
	for i = 0; i < arr1Len; i++ {
		arr1[i] = (*input)[i+start]
	}
	for i = 0; i < arr2Len; i++ {
		arr2[i] = (*input)[i+mid+1]
	}
	i, j = 0, 0
	for i < arr1Len && j < arr2Len {
		if arr1[i] < arr2[j] {
			(*input)[start] = arr1[i]
			i++
		} else {
			(*input)[start] = arr2[j]
			j++
		}
		start++
	}
	for i < arr1Len {
		(*input)[start] = arr1[i]
		i++
		start++
	}
	for j < arr2Len {
		(*input)[start] = arr2[j]
		j++
	}
	start++
}
