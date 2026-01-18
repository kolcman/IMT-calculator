package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const IMT_POWER = 2

var result float64

func main() {
	for {
		showMenu()
		input := getChoice()
		if input == "q" {
			break
		}
		result, err := calculateIMT(getUserInfo())
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(result)
		printResult(result)
		fmt.Println("******************************")
	}
}

func showMenu() {
	fmt.Println("************** Калькулятор ИМТ ***************")
	fmt.Println("     Для продложения нажмите любую кнопку		 ")
	fmt.Println("     \"Q\" - Выход из программы							 ")
	fmt.Println("**********************************************")
}

func getChoice() string {
	var input string
	fmt.Scan(&input)
	inputString := strings.TrimSpace(strings.ToLower(input))
	return inputString
}

func getUserInfo() (float64, float64) {
	var height, weight float64
	fmt.Printf("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)
	fmt.Printf("Введите свой вес в килограммах: ")
	fmt.Scan(&weight)
	return height, weight
}

func calculateIMT(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("Ошибка! Не указан вес или рост")
	}
	IMT := weight / math.Pow(height/100, IMT_POWER)
	return IMT, nil
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
