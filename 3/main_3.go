package main

import "fmt"

// StringIntMap структура для хранения пар "строка - целое число"
type StringIntMap struct {
	data map[string]int
}

// Добавление элемента
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Удаление элемента
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Копирование карты
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}

// Проверка наличия ключа
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Получение значения
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}
func main() {
	// Пример использования StringIntMap
	sim := StringIntMap{data: make(map[string]int)}
	sim.Add("один", 1)
	sim.Add("два", 2)
	fmt.Println("Мапа после добавления элементов:", sim.data)
	if sim.Exists("один") {
		value, exists := sim.Get("один") // Получаем оба значения
		if exists {
			fmt.Println("Ключь 'один' существует созначением:", value)
		}
	}
	sim.Remove("один")
	fmt.Println("Мапа после удаления ключа 'один':", sim.data)
	copiedMap := sim.Copy()
	fmt.Println("Скопированная мапа:", copiedMap)
}
