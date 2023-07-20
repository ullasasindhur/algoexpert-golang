package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var input string = get_input(scanner)
	var stringArray *[]string = split_string(&input, ' ')
	var intArray *[]int = stringArray_to_intArray(stringArray)
	bubble_sort(intArray)
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

func get_input(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func stringArray_to_intArray(stringArray *[]string) *[]int {
	var intArray []int = make([]int, 0)
	for i := 0; i < len(*stringArray); i++ {
		value, err := strconv.Atoi((*stringArray)[i])
		if err != nil {
			fmt.Println("Given input contains invalid data type or entered a trialing space")
			os.Exit(0)
		} else {
			intArray = append(intArray, value)
		}
	}
	return &intArray
}

func bubble_sort(array *[]int) {
	var arrayLen int = len(*array)
	var counter int
	var isSorted bool
	for !isSorted {
		isSorted = true
		for i := 0; i < arrayLen-1-counter; i++ {
			if (*array)[i] > (*array)[i+1] {
				(*array)[i], (*array)[i+1] = (*array)[i+1], (*array)[i]
				isSorted = false
			}
		}
		counter++
	}
	fmt.Println(*array)
}
