package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация слайса целых чисел
func generateRandomSlice(size int) []int {
	// Создаем новый генератор случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = r.Intn(100) // Генерируем случайные числа от 0 до 99
	}
	return slice
}

// Функция, возвращающая новый слайс с четными числами
func sliceExample(originalSlice []int) []int {
	var evenSlice []int
	for _, v := range originalSlice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

// Функция, добавляющая элемент в конец слайса
func addElements(slice []int, element int) []int {
	return append(slice, element)
}

// Функция, копирующая слайс
func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// Функция, удаляющая элемент по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // Возвращаем оригинальный слайс, если индекс вне диапазона
	}
	return append(slice[:index], slice[index+1:]...)
}
func main() {
	// Генерируем слайс целых чисел
	originalSlice := generateRandomSlice(10)
	fmt.Println("Исходный слайс:", originalSlice)
	// Пример работы с функцией sliceExample
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Четные числа из исходного слайса:", evenSlice)
	// Пример работы с функцией addElements
	newSlice := addElements(originalSlice, 42)
	fmt.Println("Слайс после добавления элемента 42:", newSlice)
	// Пример работы с функцией copySlice
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия исходного слайса:", copiedSlice)
	// Пример работы с функцией removeElement
	removedSlice := removeElement(originalSlice, 2) // Удаляем элемент по индексу 2
	fmt.Println("Слайс после удаления элемента по индексу 2:", removedSlice)
}
