package main

import (
	"fmt"
	"lasts-lab/internal/input"
	"slices"
)

func main() {
	// Ввод массива K.
	fmt.Print("Размер массива K> ")
	var arrSizeK uint
	_, err := fmt.Scanln(&arrSizeK)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	if arrSizeK < 0 {
		fmt.Println("Размер массива должен быть положительным")
		return
	}

	arrK, err := input.Input(arrSizeK)
	if err != nil {
		fmt.Println("Ошибка ввода: ", err.Error())
		return
	}

	// Ввод массива L
	fmt.Print("Размер массива L> ")
	var arrSizeL uint
	_, err = fmt.Scanln(&arrSizeL)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	if arrSizeL < 0 {
		fmt.Println("Размер массива должен быть положительным")
		return
	}

	arrL, err := input.Input(arrSizeL)
	if err != nil {
		fmt.Println("Ошибка ввода: ", err.Error())
		return
	}

	// Поиск наибольшего элемента в K, которого нет в L.
	maxVal := -1
	for _, val := range arrK {
		if !slices.Contains(arrL, val) && val > maxVal {
			maxVal = val
		}
	}

	if maxVal == -1 {
		fmt.Println("Нет элементов в K, которые отсутствуют в L")
	} else {
		fmt.Println("Наибольший элемент в K, которого нет в L: ", maxVal)
	}
}
