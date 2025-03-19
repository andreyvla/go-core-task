package main

import (
	"reflect"
	"testing"
)

// Тест для метода Add
func TestAdd(t *testing.T) {
	sim := StringIntMap{data: make(map[string]int)}
	sim.Add("один", 1)
	expected := 1
	if sim.data["один"] != expected {
		t.Errorf("Ожидалось значение %d, получено %d", expected, sim.data["один"])
	}
}

// Тест для метода Remove
func TestRemove(t *testing.T) {
	sim := StringIntMap{data: make(map[string]int)}
	sim.Add("два", 2)
	sim.Remove("два")
	if sim.Exists("два") {
		t.Error("Ключ 'два' должен быть удален, но он все еще существует")
	}
}

// Тест для метода Copy
func TestCopy(t *testing.T) {
	sim := StringIntMap{data: make(map[string]int)}
	sim.Add("три", 3)
	copiedMap := sim.Copy()
	// Проверяем, что копия равна оригиналу
	if !reflect.DeepEqual(sim.data, copiedMap) {
		t.Error("Скопированная мапа должна быть равна оригинальной мапе")
	}
	// Изменяем оригинальную мапу и проверяем, что копия не изменилась
	sim.Add("четыре", 4)
	if reflect.DeepEqual(sim.data, copiedMap) {
		t.Error("Скопированная мапа не должна измениться после изменения оригинала")
	}
}

// Тест для метода Exists
func TestExists(t *testing.T) {
	sim := StringIntMap{data: make(map[string]int)}
	sim.Add("пять", 5)
	if !sim.Exists("пять") {
		t.Error("Ключ 'пять' должен существовать, но он не найден")
	}
	if sim.Exists("шесть") {
		t.Error("Ключ 'шесть' не должен существовать, но он найден")
	}
}

// Тест для метода Get
func TestGet(t *testing.T) {
	sim := StringIntMap{data: make(map[string]int)}
	sim.Add("семь", 7)
	value, exists := sim.Get("семь")
	if !exists {
		t.Error("Ключ 'семь' должен существовать, но он не найден")
	}
	if value != 7 {
		t.Errorf("Ожидалось значение 7, получено %d", value)
	}
	// Проверка на несуществующий ключ
	_, exists = sim.Get("восемь")
	if exists {
		t.Error("Ключ 'восемь' не должен существовать, но он найден")
	}
}
