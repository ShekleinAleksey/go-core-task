package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestConvertToString(t *testing.T) {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	result := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	expected := []string{"42", "42", "42", "3.14", "Golang", "true", "(1+2i)"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestConcatenateStrings(t *testing.T) {
	testSlice := []string{"42", "42", "42", "3.14", "Golang", "true", "(1+2i)"}
	result := concatenateStrings(testSlice)

	expected := "42 42 42 3.14 Golang true (1+2i)"

	if result != expected {
		t.Errorf("Ожидалось '%s', получено '%s'", expected, result)
	}
}

func TestStringToRunes(t *testing.T) {
	testString := "Привет, мир"
	result := stringToRunes(testString)

	expected := []rune{'П', 'р', 'и', 'в', 'е', 'т', ',', ' ', 'м', 'и', 'р'}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestHashRunes(t *testing.T) {
	testRunes := []rune("test string")
	result := hashRunes(testRunes)

	if len(result) != 64 {
		t.Errorf("Ожидалась строка длиной 64 символа, получено %d", len(result))
	}

	for _, char := range result {
		if !strings.Contains("0123456789abcdef", string(char)) {
			t.Errorf("Недопустимый символ в хеше: %c", char)
		}
	}

	result2 := hashRunes(testRunes)
	if result != result2 {
		t.Errorf("Хеши для одинаковых входных данных не совпадают")
	}
}
