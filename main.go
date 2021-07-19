package main

import "fmt"

func main() {
	fmt.Println("==Insertion Sort==")
	input := []int{5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted algo : ", input)
	input = insertionSort(input)
	fmt.Println("Sorted algo : ", input)
}

func insertionSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		key := input[i]
		for j := i - 1; j >= 0; j-- {
			if key < input[j] {
				input[j+1], input[j] = input[j], key
			}
		}
	}
	return input
}
