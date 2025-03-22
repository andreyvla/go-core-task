package main

import (
	"math"
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	// Запускаем конвейер.
	pipeline(in, out)

	// Отправляем числа в канал in.
	go func() {
		defer close(in)
		for i := uint8(1); i <= 3; i++ {
			in <- i
		}
	}()

	// Собираем результаты из канала out.
	var results []float64
	for result := range out {
		results = append(results, result)
	}

	// Проверяем результаты.
	expected := []float64{1, 8, 27}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("ожидали %v, получили %v", expected, results)
	}
}

func TestPipelineEmptyChannel(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	close(in)

	pipeline(in, out)

	_, ok := <-out
	if ok {
		t.Error("Ожидается закрытый канал, а имеем открытый")
	}
}

func TestPipelineMultipleNumbers(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	pipeline(in, out)

	go func() {
		defer close(in)
		for i := uint8(1); i <= 10; i++ {
			in <- i
		}
	}()

	var results []float64
	for result := range out {
		results = append(results, result)
	}

	expected := make([]float64, 10)
	for i := 0; i < 10; i++ {
		expected[i] = math.Pow(float64(i+1), 3)
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("ожидали %v, получили %v", expected, results)
	}
}
