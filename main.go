package main

import (
	"fmt"
)

func main() {
	A := []string{"cat", "cat", "dog", "cat", "tree"}
	checker := make(map[string]struct{})
	NewAsset := make([]string, 0)
	for _, v := range A {
		_, ok := checker[v]
		if !ok {
			NewAsset = append(NewAsset, v)
			checker[v] = struct{}{}
		}
	}

	fmt.Println(NewAsset)
}
