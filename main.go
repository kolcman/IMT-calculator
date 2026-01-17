package main

import (
	"fmt"
	"math"
)

func main() {
	const IMT_POWER = 2
	height, weight := 1.8, 100.0
	IMT := weight / math.Pow(height, IMT_POWER)
	fmt.Printf("Yout IMT: %f", IMT)
}
