package main

import (
	"strings"
)

var justString string //предположим что эта переменная должна быть глобальной

func main() {

	someFunc()
	justString = someFunc2()
}

func someFunc2() string { // этот вариант так-же сделает присваивание значения переменной justString
	v := createHugeString(1 << 10) // более наглядным, а функцию универсальной
	return strings.Clone(v[:100])
}

func someFunc() {
	v := createHugeString(1 << 10) //justString = v[:100]  в данной реализации justString ссылается на небольшой фрагмент v
	newV := strings.Clone(v[:100]) // тем самым не давая сборщику мусора очистить память от лишних символов
	justString = newV              // эта строка кода решает проблему, создаваю новую строку
} // меньшего размера, тем самым позволяя сборщику очистить память от огромной v
