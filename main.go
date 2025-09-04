package main

import (
	"fmt"
	"math"
)

func main() {
	orderedTemp := make(map[string][]string)
	negZero := math.Copysign(0, -1)
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5, 5}
	for _, v := range temp {
		if v < 0 && v > -10 {
			orderedTemp[fmt.Sprintf("%.0f", negZero)] = append(orderedTemp[fmt.Sprintf("%.0f", negZero)], fmt.Sprintf("%.1f", v))
			continue
		}
		orderedTemp[fmt.Sprintf("%.0f", math.Floor(v/10)*10)] = append(orderedTemp[fmt.Sprintf("%.0f", math.Floor(v/10)*10)], fmt.Sprintf("%.1f", v))
		// orderedTemp[int(v/10)*10] = append(orderedTemp[int(v/10)*10], v) в случае если можно пренебречь отображением 0 после запятой
	}
	fmt.Println(orderedTemp)
}
