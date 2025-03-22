package main

import (
	"fmt"
)

// findIntersection проверяет, есть ли пересечения между двумя слайсами int и возвращает:
// - bool: true, если есть хотя бы одно пересечение, false в противном случае.
// - []int: срез с пересекающимися значениями (пустой срез, если пересечений нет).
func findIntersection(slice1, slice2 []int) (bool, []int) {
	// Создаем мапу для быстрого поиска элементов из первого слайса.
	map1 := make(map[int]bool)
	for _, num := range slice1 {
		map1[num] = true
	}

	// Создаем срез для хранения пересекающихся значений.
	intersection := []int{}
	// Флаг, указывающий на наличие пересечений.
	hasIntersection := false

	// Проходим по второму слайсу и проверяем, есть ли элементы в мапе.
	for _, num := range slice2 {
		if map1[num] {
			// Если элемент есть в мапе, добавляем его в срез пересечений.
			// Проверяем, не добавлен ли он уже.
			found := false
			for _, val := range intersection {
				if val == num {
					found = true
					break
				}
			}
			if !found {
				intersection = append(intersection, num)
			}

			// Устанавливаем флаг наличия пересечений в true.
			hasIntersection = true
		}
	}

	return hasIntersection, intersection
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersection := findIntersection(a, b)
	fmt.Println("Слайс a:", a)
	fmt.Println("Слайс b:", b)
	fmt.Println("Есть пересечения:", hasIntersection)
	fmt.Println("Пересечения:", intersection)

	c := []int{1, 2, 3}
	d := []int{4, 5, 6}
	hasIntersection, intersection = findIntersection(c, d)
	fmt.Println("Слайс c:", c)
	fmt.Println("Слайс d:", d)
	fmt.Println("Есть пересечения:", hasIntersection)
	fmt.Println("Пересечения:", intersection)

	e := []int{1, 2, 3, 3, 2, 1}
	f := []int{3, 2, 1, 1, 2, 3}
	hasIntersection, intersection = findIntersection(e, f)
	fmt.Println("Слайс e:", e)
	fmt.Println("Слайс f:", f)
	fmt.Println("Есть пересечения:", hasIntersection)
	fmt.Println("Пересечения:", intersection)
}
