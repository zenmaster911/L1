package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{}
	slices.Sort(nums)
	fmt.Println(nums, binarySearch(1, nums))
}

func binarySearch(target int, input []int) int {
	if len(input) == 0 {
		return -1
	}
	left := 0
	right := len(input) - 1
	i := (left + right + 1) / 2
	for {
		if target < input[left] || target > input[right] {
			return -1
		}
		if input[i] == target {
			return i
		} else if input[i] < target {
			left = i + 1
			i = (left + right + 1) / 2
		} else {
			right = i - 1
			i = (left + right) / 2
		}

	}
}
