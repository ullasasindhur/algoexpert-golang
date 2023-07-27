package count_inversions

import "fmt"

func Count_inversions(input *[]int) {
	fmt.Print("Count Inversions: ")
	var inversions int
	split(input, 0, len(*input)-1, &inversions)
	fmt.Println(inversions)
}

func split(input *[]int, start int, end int, inversions *int) {
	if start < end {
		var mid int = (start + end) / 2
		split(input, start, mid, inversions)
		split(input, mid+1, end, inversions)
		count(input, start, mid, end, inversions)
	}
}

func count(input *[]int, start int, mid int, end int, inversions *int) {
	var arr1Len, arr2Len int = mid - start + 1, end - mid
	var arr1, arr2 []int = make([]int, arr1Len), make([]int, arr2Len)
	var i, j int
	for i = 0; i < arr1Len; i++ {
		arr1[i] = (*input)[start+i]
	}
	for j = 0; j < arr2Len; j++ {
		arr2[j] = (*input)[mid+j+1]
	}
	i, j = 0, 0
	for i < arr1Len && j < arr2Len {
		if arr1[i] < arr2[j] {
			(*input)[start] = arr1[i]
			i++
		} else {
			(*input)[start] = arr2[j]
			j++
			*inversions++
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
		start++
	}
}
