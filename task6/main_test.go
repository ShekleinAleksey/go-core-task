package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomNumGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := RandomNumGenerator(ctx)

	select {
	case val := <-ch:
		if val < 0 || val > 99 {
			t.Errorf("ожидалось число от 0 до 100, получено %d", val)
		}
	case <-time.After(time.Second):
		t.Error("принимающий канал не получил значение")
	}
}
