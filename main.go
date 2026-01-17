package main

import (
	"fmt"
	"math"
)

const IMT_POWER = 2

var result float64

func main() {
	fmt.Println("******* Калькулято ИМТ *******")
	result := calculateIMT(getUserIfnfo())
	fmt.Printf("Ваш ИМТ: %.1f\n", result)
	fmt.Println("******* ************** *******")
}

func getUserIfnfo() (float64, float64) {
	var height, weight float64
	fmt.Printf("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)
	fmt.Printf("Введите свой вес в киллограммах: ")
	fmt.Scan(&weight)
	return height, weight
}

func calculateIMT(height, weight float64) float64 {
	IMT := weight / math.Pow(height/100, IMT_POWER)
	return IMT
}
