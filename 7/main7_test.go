package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	// Создаем несколько каналов.
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	// Отправляем данные в каналы.
	go func() {
		defer close(c1)
		for i := 1; i <= 3; i++ {
			c1 <- i
		}
	}()

	go func() {
		defer close(c2)
		for i := 4; i <= 6; i++ {
			c2 <- i
		}
	}()

	go func() {
		defer close(c3)
		for i := 7; i <= 9; i++ {
			c3 <- i
		}
	}()

	// Сливаем каналы.
	merged := merge(c1, c2, c3)

	// Собираем данные из объединенного канала.
	var result []int
	for n := range merged {
		result = append(result, n)
	}

	// Сортируем результат для сравнения.
	sort.Ints(result)

	// Проверяем, что все данные из входных каналов попали в выходной.
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидали %v, получили %v", expected, result)
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	// Создаем пустые каналы.
	c1 := make(chan int)
	c2 := make(chan int)
	close(c1)
	close(c2)

	// Сливаем каналы.
	merged := merge(c1, c2)

	// Проверяем, что канал закрыт.
	_, ok := <-merged
	if ok {
		t.Error("Ожидали закрытый канал, а получили открытый")
	}
}
