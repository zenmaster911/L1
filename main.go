package main

import (
	"fmt"
)

func main() {
	oldSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	index := 3
	newSlice1 := shorter1(oldSlice, index)
	newSlice2 := shorter2(oldSlice, index)
	fmt.Println(newSlice1, newSlice2)
}

func shorter1[S ~[]E, E any](slice S, ind int) S {
	newSlice := make([]E, len(slice)-1, len(slice)-1)
	copy(newSlice[:ind], slice[:ind])
	copy(newSlice[ind:], slice[ind+1:])
	return newSlice
}

func shorter2[S ~[]E, E any](slice S, ind int) S {
	newSlice := make([]E, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:ind]...)
	newSlice = append(newSlice, slice[ind+1:]...)
	return newSlice
}
