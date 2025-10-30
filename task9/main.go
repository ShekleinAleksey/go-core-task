// Сделать конвейер чисел Даны два канала. В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления, затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.

// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

// Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
)

func main() {
	senderChan := make(chan uint8)
	reciveChan := make(chan float64)

	go func() {
		numbers := []uint8{1, 2, 3, 4, 5}
		for _, val := range numbers {
			senderChan <- val
		}
		close(senderChan)
	}()

	go CalcPipeline(senderChan, reciveChan)

	for val := range reciveChan {
		fmt.Println(val)
	}

}

func CalcPipeline(senderChan chan uint8, reciveChan chan float64) {

	for val := range senderChan {
		result := float64(val) * float64(val) * float64(val)
		reciveChan <- result
	}
	close(reciveChan)
}
