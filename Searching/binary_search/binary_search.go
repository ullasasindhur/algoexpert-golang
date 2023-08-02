package binary_search

import (
	"fmt"
	"os"
)

func Binary_search(input []int, value int) {
	fmt.Println("Binary Search: ")
	search(input, value, 0, len(input)-1)
	fmt.Println(value, "is not found in the input array")
}

func search(input []int, value int, start int, end int) {
	if start <= end {
		var mid int = (start + end) / 2
		if input[mid] == value {
			fmt.Println(value, "is found in the input array")
			os.Exit(0)
		} else if input[mid] < value {
			search(input, value, mid+1, end)
		} else {
			search(input, value, start, mid-1)
		}
	}
}
