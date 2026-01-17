package main

import (
	"fmt"
	"math"
)

const IMT_POWER = 2

var result float64

func main() {
	fmt.Println("******* Калькулято ИМТ *******")
	result := calculateIMT(readValue("рост"), readValue("вес"))
	fmt.Printf("Ваш ИМТ: %.1f\n", result)
	fmt.Println("******* ************** *******")
}

func readValue(typeValue string) float64 {
	var height float64
	fmt.Printf("Введите свой %v: ", typeValue)
	fmt.Scan(&height)
	return height
}

func calculateIMT(height, weight float64) float64 {
	IMT := weight / math.Pow(height/100, IMT_POWER)
	return IMT
}
