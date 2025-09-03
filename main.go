package main

import "fmt"

func main() {
	orderedTemp := make(map[int][]string)
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, v := range temp {
		orderedTemp[int(v/10)*10] = append(orderedTemp[int(v/10)*10], fmt.Sprintf("%.1f", v))
		// orderedTemp[int(v/10)*10] = append(orderedTemp[int(v/10)*10], v) в случае если можно пренебречь отображением 0 после запятой
	}
	fmt.Println(orderedTemp)
}
