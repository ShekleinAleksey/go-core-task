package main

import (
	"reflect"
	"testing"
)

func TestElemFromFirstSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := ElemFromFirstSlice(slice1, slice2)

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestIdenticalElements(t *testing.T) {
	slice1 := []string{"apple", "banana"}
	slice2 := []string{"apple", "banana", "date"}

	result := ElemFromFirstSlice(slice1, slice2)

	expected := []string{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось пусто, получено %v", result)
	}
}
