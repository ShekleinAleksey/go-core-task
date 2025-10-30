package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, result := Intersection(a, b)

	expected := []int{64, 3}

	if !ok {
		t.Errorf("ожидалось, что пересечение есть")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestNoIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{67, 2, 11, 43}

	ok, result := Intersection(a, b)

	if ok {
		t.Errorf("ожидалось отсутствие пересечения")
	}

	if len(result) != 0 {
		t.Errorf("ожидался пустой срез, получено %v", result)
	}
}
