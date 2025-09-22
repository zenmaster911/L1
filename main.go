package main

import (
	"fmt"
	"unicode"
)

func main() {
	str1 := "asdfghjфывт"
	str2 := "asdfAghj"
	str3 := "ФФывт"
	fmt.Println(UnicRunesCheck(str1))
	fmt.Println(UnicRunesCheck(str2))
	fmt.Println(UnicRunesCheck(str3))

}

func UnicRunesCheck(str string) bool {
	checkmap := make(map[rune]struct{})
	for _, v := range str {
		val := unicode.ToLower(v)
		if _, ok := checkmap[val]; ok {
			return false
		}
		checkmap[val] = struct{}{}
	}
	return true
}
