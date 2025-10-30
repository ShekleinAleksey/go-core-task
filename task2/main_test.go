package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sliceExample(originalSlice)

	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось '%v', получено '%v'", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		value    int
		expected []int
	}{
		{
			name:     "добавление в пустой срез",
			input:    []int{},
			value:    5,
			expected: []int{5},
		},
		{
			name:     "добавление в непустой срез",
			input:    []int{1, 2, 3},
			value:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "добавление отрицательного числа",
			input:    []int{10, 20},
			value:    -5,
			expected: []int{10, 20, -5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := addElements(tt.input, tt.value)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, ожидалось %v", tt.input, tt.value, result, tt.expected)
			}

			if !reflect.DeepEqual(tt.input, original) {
				t.Errorf("исходный срез изменился: было %v, стало %v", original, tt.input)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "копирование обычного слайса",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "копирование пустого слайса",
			input: []int{},
		},
		{
			name:  "копирование слайса с одним элементом",
			input: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)

			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("copySlice(%v) = %v, ожидалось %v", tt.input, result, tt.input)
			}

			if len(tt.input) > 0 {
				result[0] = result[0] + 100
				if result[0] == tt.input[0] {
					t.Errorf("копирование работает не корректно")
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		index       int
		expected    []int
		shouldPanic bool
	}{
		{
			name:     "удаление из середины",
			input:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "удаление первого элемента",
			input:    []int{1, 2, 3},
			index:    0,
			expected: []int{2, 3},
		},
		{
			name:     "удаление последнего элемента",
			input:    []int{1, 2, 3, 4},
			index:    3,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := removeElement(tt.input, tt.index)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("removeElement(%v, %d) = %v, ожидалось %v", tt.input, tt.index, result, tt.expected)
			}

			if !reflect.DeepEqual(tt.input, original) {
				t.Errorf("исходный срез изменился: было %v, стало %v", original, tt.input)
			}
		})
	}
}
