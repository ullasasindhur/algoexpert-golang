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
	insertion_sort(intArray)
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
			fmt.Println("Error in input data type or trailing space is present in input.")
			os.Exit(0)
		} else {
			intArray = append(intArray, value)
		}
	}
	return &intArray
}

func insertion_sort(array *[]int) {
	var arrayLen int = len(*array)
	var i, j int
	for i = 1; i < arrayLen; i++ {
		j = i
		for j > 0 && (*array)[j-1] > (*array)[j] {
			(*array)[j-1], (*array)[j] = (*array)[j], (*array)[j-1]
			j--
		}
	}
	fmt.Println(*array)
}
