package radixsort

import "fmt"

func Radix_Sort(input *[]int) {
	fmt.Println("Radix Sort:")
	var maxNumber *int = get_max_number(*input)
	var digitPlace int = 1
	for *maxNumber > 0 {
		radix(input, digitPlace)
		*maxNumber /= 10
		digitPlace *= 10
	}
	fmt.Println(*input)
}

func get_max_number(input []int) *int {
	var max int = 0
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return &max
}

func radix(input *[]int, digitPlace int) {
	var inputLen int = len(*input)
	var countArray []int = make([]int, 10)
	var tempArray []int = make([]int, inputLen)
	var index int
	for i := 0; i < inputLen; i++ {
		countArray[((*input)[i]/digitPlace)%10] += 1
	}
	for i := 1; i < 10; i++ {
		countArray[i] += countArray[i-1]
	}
	for i := inputLen - 1; i >= 0; i-- {
		index = ((*input)[i] / digitPlace) % 10
		countArray[index] -= 1
		tempArray[countArray[index]] = (*input)[i]
	}
	for i := 0; i < inputLen; i++ {
		(*input)[i] = tempArray[i]
	}
}
