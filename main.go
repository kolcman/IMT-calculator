package main

import (
	"fmt"
	"math"
)

const IMT_POWER = 2

var result float64

func main() {
	fmt.Println("******* Калькулято ИМТ *******")
	result := calculateIMT(getUserInfo())
	fmt.Print(result)
	printResult(result)
	fmt.Println("******* ************** *******")
}

func getUserInfo() (float64, float64) {
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

func printResult(imt float64) {
	fmt.Printf("Ваш ИМТ: %.1f\n", imt)
	switch {
	case imt < 16:
		fmt.Println("У Вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У Вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У Вас нормальный вес")
	case imt < 30:
		fmt.Println("У Вас избыточный вес")
	case imt < 35:
		fmt.Println("У Вас 1 степень ожирения")
	case imt < 40:
		fmt.Println("У Вас 2 степень ожирения")
	case imt >= 40:
		fmt.Println("У Вас 3 степень ожирения")
	default:
		fmt.Println("Некорректное значение ИМТ")
	}
}
