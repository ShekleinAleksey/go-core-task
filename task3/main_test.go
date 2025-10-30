package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	val, ok := m.Get("a")
	if val != 1 || !ok {
		t.Errorf("ожидалось Get('a') = 1, получено %d, ok=%v", val, ok)
	}

	val, ok = m.Get("b")
	if val != 2 || !ok {
		t.Errorf("ожидалось Get('b') = 2, получено %d, ok=%v", val, ok)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Remove("a")

	_, ok := m.Get("a")
	if ok {
		t.Errorf("элемент 'a' должен быть удален")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)

	if !m.Exists("a") {
		t.Errorf("ожидалось, что ключ 'a' существует")
	}

	if m.Exists("b") {
		t.Errorf("ожидалось, что ключ 'b' не существует")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)

	copyMap := m.Copy()
	if !reflect.DeepEqual(copyMap, m.data) {
		t.Errorf("копия не совпадает с оригиналом: %v vs %v", copyMap, m.data)
	}

	copyMap["a"] = 10
	if m.data["a"] == 10 {
		t.Errorf("копия должна быть независимой от оригинала")
	}
}
