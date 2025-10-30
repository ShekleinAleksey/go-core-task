// Напишите программу на Go, которая сливает N каналов в один.

// Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		a <- 1
		close(a)
	}()

	go func() {
		b <- 2
		close(b)
	}()

	go func() {
		c <- 3
		close(c)
	}()

	mergedChan := mergeChan(a, b, c)
	for num := range mergedChan {
		fmt.Println(num)
	}

}

func mergeChan(chans ...<-chan int) <-chan int {
	mergedChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chans))

	// Запускаем горутину для каждого входного канала
	for _, ch := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for value := range c {
				mergedChan <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()
	return mergedChan
}
