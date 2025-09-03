package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	checker := make(map[int]struct{})
	intersection := make([]int, 0)
	for _, v := range A {
		checker[v] = struct{}{}
	}
	for _, v := range B {
		_, ok := checker[v]
		if ok {
			intersection = append(intersection, v)
		}
	}
	fmt.Println(intersection)
}
