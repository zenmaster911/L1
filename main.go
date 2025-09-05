package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 6, 3, 5, 76, 45, 64, -92}
	fmt.Println(quickSort(nums))
}

func quickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	pivot := input[0]
	left := make([]int, 0)
	right := make([]int, 0)
	middle := make([]int, 0)
	for _, v := range input {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	left = append(left, middle...)
	left = append(left, right...)
	return left

}
