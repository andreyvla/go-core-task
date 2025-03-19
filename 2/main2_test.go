package main

import (
	"reflect"
	"testing"
)

// Тест для функции generateRandomSlice
func TestGenerateRandomSlice(t *testing.T) {
	slice := generateRandomSlice(10)
	if len(slice) != 10 {
		t.Errorf("Ожидалось 10 элементов в слайсе, но получили %d", len(slice))
	}
	for _, v := range slice {
		if v < 0 || v >= 100 {
			t.Errorf("Ожидали значения от 0 до 99, но получили %d", v)
		}
	}
}

// Тест для функции sliceExample
func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSlice := []int{2, 4, 6, 8, 10}
	result := sliceExample(originalSlice)
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидали %v, получили %v", expectedSlice, result)
	}
}

// Тест для функции addElements
func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	element := 4
	expectedSlice := []int{1, 2, 3, 4}
	result := addElements(slice, element)
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидали %v, получили %v", expectedSlice, result)
	}
}

// Тест для функции copySlice
func TestCopySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	copiedSlice := copySlice(originalSlice)
	if !reflect.DeepEqual(originalSlice, copiedSlice) {
		t.Errorf("Ожидается, что копия слайса будет совпадать с оригиналом")
	}
	// Изменяем оригинальный слайс и проверяем, что копия не изменилась
	originalSlice[0] = 99
	if reflect.DeepEqual(originalSlice, copiedSlice) {
		t.Errorf("Ожидается, что копия слайса будет отличаться от оригинала")
	}
}

// Тест для функции removeElement
func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expectedSlice := []int{1, 2, 4, 5}
	result := removeElement(slice, 2) // Удаляем элемент по индексу 2
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидали %v, получили %v", expectedSlice, result)
	}
	// Проверка на удаление элемента с некорректным индексом
	result = removeElement(slice, 10) // Удаляем элемент по некорректному индексу
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("Ожидали %v, получили %v", slice, result)
	}
}
