package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("******* Калькулято ИМТ *******")
	const IMT_POWER = 2
	var height, weight float64
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&height)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&weight)
	IMT := weight / math.Pow(height, IMT_POWER)
	fmt.Printf("Ваш ИМТ: %f", IMT)
}
