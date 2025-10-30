// Задание 2
// Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным образом (при каждом запуске должны получаться новые значения)

// Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.

// Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.

// Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.

// Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.

// Напишите main функцию, в которой протестируете все вышеописанные функции. Выведите результаты на экран.

// Напишите unit тесты к созданным функциям
// Примечание. В качестве originalSlice можно использовать originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var originalSlice = make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(100))
	}
	fmt.Println(originalSlice)
	evenNum := sliceExample(originalSlice)
	fmt.Println(evenNum)

	addSlice := addElements(originalSlice, 5)
	fmt.Println(addSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println(copiedSlice)

	removedSlice := removeElement(originalSlice, 4)
	fmt.Println(removedSlice)
}

func sliceExample(originalSlice []int) []int {
	var evenNum []int
	for _, val := range originalSlice {
		if val%2 == 0 {
			evenNum = append(evenNum, val)
		}
	}

	return evenNum
}

func addElements(origSlice []int, val int) []int {
	origSlice = append(origSlice, val)

	return origSlice
}

func copySlice(sliceForCopy []int) []int {
	copiedSlice := make([]int, len(sliceForCopy))
	copy(copiedSlice, sliceForCopy)

	return copiedSlice
}

func removeElement(sliceForRemove []int, i int) []int {
	if i < 0 || i >= len(sliceForRemove) {
		return sliceForRemove
	}

	result := make([]int, 0, len(sliceForRemove)-1)
	result = append(result, sliceForRemove[:i]...)
	result = append(result, sliceForRemove[i+1:]...)
	return result
}
