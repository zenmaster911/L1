package main

import (
	"fmt"
	"math/big"
)

func main() {
	for {

		var a, b big.Int
		var operator string
		var stop string
		fmt.Println("желаете продолжить? y/n")
		fmt.Scan(&stop)
		if stop == "n" {
			break
		} else if stop != "y" {
			fmt.Println("команда неясна, попробуйте снова")
		}
		fmt.Println("Введите первое число")

		if _, err := fmt.Scan(&a); err != nil {
			fmt.Printf("first number scanning error: %v\n", err)
			continue
		}

		fmt.Println("Введите один из следующих операндов: '+' '-' '*' '/'")
		if _, err := fmt.Scan(&operator); err != nil {
			//fmt.Errorf()
			fmt.Printf("operator scanning error: %v\n", err)
			continue
		}

		fmt.Println("Введите второе число")
		if _, err := fmt.Scanln(&b); err != nil {
			fmt.Printf("second number scanning error: %v", err)
			continue
		}

		switch operator {
		case "+":
			fmt.Printf("результат операции: %v\n", a.Add(&a, &b))
		case "*":
			fmt.Printf("результат операции: %v\n", a.Mul(&a, &b))
		case "-":
			fmt.Printf("результат операции: %v\n", a.Sub(&a, &b))
		case "/":
			fmt.Printf("результат операции: %v\n", a.Div(&a, &b))
		default:
			fmt.Println("invalid operator")
		}

	}
}
