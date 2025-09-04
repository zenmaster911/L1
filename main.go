package main

import (
	"fmt"
)

func main() {

	fmt.Println(mathSwap(5, 10))
	fmt.Println(bitSwap(5, 10))

}

func mathSwap(a, b int) (newA, newB int) {
	a += b
	b = a - b
	a -= b
	return a, b
}

func bitSwap(a, b int) (newA, newB int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}
