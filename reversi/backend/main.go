package main

import (
	"fmt"
	"math"
)

func shake(x float64, y float64) float64 {
	return math.Sin(x) * math.Sin(y)
}

func main() {
	fmt.Println(shake(10, 5))
}
