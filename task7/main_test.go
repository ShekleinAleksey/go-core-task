package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge_Basic(t *testing.T) {
	a := make(chan int, 2)
	b := make(chan int, 2)
	c := make(chan int, 2)

	a <- 1
	a <- 2
	close(a)

	b <- 3
	b <- 4
	close(b)

	c <- 5
	c <- 6
	close(c)

	mergeChan := mergeChan(a, b, c)

	var result []int
	for val := range mergeChan {
		result = append(result, val)
	}

	sort.Ints(result)
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}
