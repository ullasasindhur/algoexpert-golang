package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var input = get_input(scanner)
	var stringArray *[]string = split_string(&input, ' ')
	var intArray *[]int = stringArray_to_intArray(stringArray)
	selection_sort(intArray)
}

func get_input(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func split_string(input *string, delim byte) *[]string {
	var stringArray []string = make([]string, 0)
	var index int
	for i := 0; i < len(*input); i++ {
		if (*input)[i] == delim {
			stringArray = append(stringArray, (*input)[index:i])
			index = i + 1
		}
	}
	stringArray = append(stringArray, (*input)[index:])
	return &stringArray
}

func stringArray_to_intArray(stringArray *[]string) *[]int {
	var intArray []int = make([]int, 0)
	for i := 0; i < len(*stringArray); i++ {
		value, err := strconv.Atoi((*stringArray)[i])
		if err != nil {
			fmt.Println("Invalid input data type or trailing space is present.")
			os.Exit(0)
		} else {
			intArray = append(intArray, value)
		}
	}
	return &intArray
}

func selection_sort(input *[]int) {
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
