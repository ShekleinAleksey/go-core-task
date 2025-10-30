package main

import (
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	wg := NewSemaphoreWaitGroup()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// ok
	case <-time.After(2 * time.Second):
		t.Fatal("Wait() не сработал")
	}
}
