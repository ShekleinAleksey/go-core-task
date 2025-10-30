// Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

// Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SemaphoreWaitGroup struct {
	counter int32
	ch      chan struct{}
}

func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		ch: make(chan struct{}),
	}
}

func main() {
	wg := NewSemaphoreWaitGroup()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(500 * time.Millisecond) //Имитация работы
			fmt.Println("Горутина №", id, "завершилась")
		}(i)
	}

	fmt.Println("Ожидание завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершились")
}

func (wg *SemaphoreWaitGroup) Add(delta int) {
	newCounter := atomic.AddInt32(&wg.counter, int32(delta))

	if newCounter < 0 {
		panic("Отрицательное количество горутин")
	}

	if newCounter == 0 {
		close(wg.ch)
	}
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *SemaphoreWaitGroup) Wait() {
	<-wg.ch
}
