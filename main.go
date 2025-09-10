package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("choose func you would like to use:\n print 1 to use func with Slices\n print 2 to avoid using slices")
	var choise int
	fmt.Scan(&choise)
	switch choise {
	case 1:
		withSliceslib()
	case 2:
		withoutSliceslib()
	}
}

func withSliceslib() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	fmt.Println("to stop scanning print \"end\"")
	for scanner.Scan() {
		initialPhrase := scanner.Text()
		if initialPhrase == "end" {
			break
		}
		slice := strings.Split(initialPhrase, " ")
		slices.Reverse(slice)
		newPhrase := strings.Join(slice, " ")
		fmt.Println(newPhrase)
	}
}

func withoutSliceslib() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	fmt.Println("to stop scanning print \"end\"")

	for scanner.Scan() {

		inputData := scanner.Text()

		if inputData == "end" {
			break
		}

		fmt.Println(wordOrderReverse(inputData))

	}
}

func wordOrderReverse(str string) string {
	var newstr []byte
	var end int
	end = len(str)
	for b := end - 1; b >= 0; b-- {
		if str[b] == byte(' ') {
			newstr = append(newstr, str[b:end]...)
			end = b
		}
		if b == 0 {
			newstr = append(newstr, byte(' '))
			newstr = append(newstr, str[b:end]...)
		}
	}
	newstr = newstr[1:]

	return string(newstr)
}
