package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	bubble_sort "github.com/ullasasindhur/algoexpert-golang/Sorting/bubble_sort"
	heapsort "github.com/ullasasindhur/algoexpert-golang/Sorting/heap_sort"
	insertion_sort "github.com/ullasasindhur/algoexpert-golang/Sorting/insertion_sort"
	"github.com/ullasasindhur/algoexpert-golang/Sorting/merge_sort"
	quicksort "github.com/ullasasindhur/algoexpert-golang/Sorting/quick_sort"
	radixsort "github.com/ullasasindhur/algoexpert-golang/Sorting/radix_sort"
	selection_sort "github.com/ullasasindhur/algoexpert-golang/Sorting/selection_sort"
	threenumbersort "github.com/ullasasindhur/algoexpert-golang/Sorting/three_number_sort"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var input = get_input(scanner)
	var stringArray *[]string = split_string(&input, ' ')
	var intArray *[]int = stringArray_to_intArray(stringArray)
	selection_sort.Selection_sort(get_copy_slice(*intArray))
	insertion_sort.Insertion_sort(get_copy_slice(*intArray))
	bubble_sort.Bubble_sort(get_copy_slice(*intArray))
	quicksort.Quick_sort(get_copy_slice(*intArray))
	heapsort.Heap_Sort(get_copy_slice(*intArray))
	var three_numbers_arr []int = []int{-1, 1, 0, 0, -1, 0, -1, 1, 0, -1}
	var three_number_order []int = []int{-1, 0, 1}
	threenumbersort.Three_Number_Sort(&three_numbers_arr, &three_number_order)
	radixsort.Radix_Sort(get_copy_slice(*intArray))
	merge_sort.Merge_sort(get_copy_slice(*intArray))
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

func get_copy_slice(input []int) *[]int {
	var output []int = make([]int, len(input))
	copy(output, input)
	return &output
}
