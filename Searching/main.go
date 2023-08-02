package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/ullasasindhur/algoexpert-golang/Searching/binary_search"
)

func main() {
	fmt.Println("Searching")
	var array []int = get_random_array(10)
	fmt.Println(array)
	binary_search.Binary_search(array, 5)
}

func get_random_array(length int) []int {
	var output []int = make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = rand.Intn(20)
	}
	sort.Ints(output)
	return output
}
