// Задание 3
// Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число". Ваша структура должна поддерживать следующие операции:

// Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.

// Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.

// Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.

// Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.

// Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.

// Напишите unit тесты к созданным функциям

package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func main() {
	m := NewStringIntMap()

	m.Add("a", 1)
	m.Add("b", 2)
	fmt.Println("m.Add", m)

	m.Remove("a")
	fmt.Println("m.Remove('a')", m)

	m.Add("c", 3)
	copy := m.Copy()
	fmt.Println("m.Copy()", copy)

	fmt.Println("m.Exists('a')", m.Exists("a"))
	fmt.Println("m.Exists('b')", m.Exists("b"))

}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copiedMap := make(map[string]int, len(m.data))
	for i, val := range m.data {
		copiedMap[i] = val
	}

	return copiedMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]

	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]

	return val, ok
}
