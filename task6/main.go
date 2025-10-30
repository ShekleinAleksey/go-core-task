// Напишите генератор случайных чисел используя небуфферизированный канал.

// Напишите unit тесты к созданным функциям

package main

import (
	"context"
	"fmt"
	"math/rand"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	randChan := RandomNumGenerator(ctx)

	for i := 0; i < 20; i++ {
		fmt.Println(<-randChan)
	}

}

func RandomNumGenerator(ctx context.Context) <-chan int {
	randChan := make(chan int)

	go func() {
		defer close(randChan)
		for {
			select {
			case <-ctx.Done():
				return
			case randChan <- rand.Intn(100):
			}
		}

	}()

	return randChan
}
