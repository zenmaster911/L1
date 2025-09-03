package main

import (
	"fmt"
	"log"
)

func main() {
	var number int64
	var requiredBit int
	value := true
	fmt.Println("enter the number, required bit and Value you would like to set")
	_, err := fmt.Scan(&number, &requiredBit, &value)
	if err != nil {
		log.Fatalf("data input error: %v", err)
	}
	fmt.Println("\n", bitConverser(number, requiredBit-1, value))
}

func bitConverser(number int64, bit int, value bool) int64 {
	mask := int64(1 << bit)
	if value {
		return number | mask
	}
	return number &^ mask

}
