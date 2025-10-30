// Задание 4
// На вход подаются два неупорядоченных слайса строк. Например:

// slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
// slice2 := []string{"banana", "date", "fig"}
// Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.

// Напишите unit тесты к созданным функциям

package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := ElemFromFirstSlice(slice1, slice2)
	fmt.Println(result)
}

func ElemFromFirstSlice(slice1 []string, slice2 []string) []string {
	resultSlice := make([]string, 0)
	map2 := make(map[string]bool, 0)
	for _, val := range slice2 {
		map2[val] = true
	}
	for _, val := range slice1 {
		if _, ok := map2[val]; ok != true {
			resultSlice = append(resultSlice, val)
		}
	}
	return resultSlice
}
