package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	fmt.Println("to stop scanning enter \"end\"")
	for scanner.Scan() {
		initialWord := scanner.Text()
		if initialWord == "end" {
			break
		}
		newWord := make([]rune, len(initialWord))
		for i, r := range initialWord {
			newWord[len(newWord)-1-i] = r
		}
		fmt.Println(string(newWord))
	}
}
