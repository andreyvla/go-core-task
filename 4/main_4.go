package main

import "fmt"

// findDifference находит элементы, которые есть в slice1, но отсутствуют в slice2.
func findDifference(slice1, slice2 []string) []string {
	// Создаем мапу для быстрого поиска элементов из slice2.
	slice2Map := make(map[string]bool)
	for _, item := range slice2 {
		slice2Map[item] = true
	}

	// Создаем слайс для хранения результата.
	result := []string{}

	// Проходим по slice1 и проверяем, есть ли элемент в slice2Map.
	for _, item := range slice1 {
		if !slice2Map[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	difference := findDifference(slice1, slice2)
	fmt.Println("Разница между слайсами:", difference)

	slice3 := []string{"a", "b", "c"}
	slice4 := []string{"d", "e", "f"}
	difference2 := findDifference(slice3, slice4)
	fmt.Println("Разница между слайсами:", difference2)

	slice5 := []string{"a", "b", "c"}
	slice6 := []string{"a", "b", "c"}
	difference3 := findDifference(slice5, slice6)
	fmt.Println("Разница между слайсами:", difference3)
}
