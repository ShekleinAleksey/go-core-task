// На вход подаются два неупорядоченных слайса int любой длины. Например:

// a := []int{65, 3, 58, 678, 64}
// b := []int{64, 2, 3, 43}
// Напишите функцию, которая проверяет, есть ли пересечения значений между двумя слайсами и возвращает:

// bool значение есть ли хотя бы одно пересечение в значениях входных срезов
// срез []int с пересеченными значениями (если таких значений нет, то возвращать пустой срез). Т. е. если взать слайсы a и b из премере, то вернет true, []int{64, 3}.
// Напишите unit тесты к созданным функциям

package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, result := Intersection(a, b)
	fmt.Println(ok, " ", result)
}

func Intersection(slice1 []int, slice2 []int) (bool, []int) {
	resultSlice := make([]int, 0)
	tempMap := make(map[int]bool, 0)
	for _, val := range slice1 {
		tempMap[val] = true
	}
	for _, val := range slice2 {
		if _, ok := tempMap[val]; ok == true {
			resultSlice = append(resultSlice, val)
		}
	}

	if len(resultSlice) != 0 {
		return true, resultSlice
	} else {
		return false, []int{}
	}
}
