package main

import (
	m "Tutorial/Chapter11/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := m.Average(xs)
	fmt.Println("Result Average:", avg)
}
