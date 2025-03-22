package main

import (
	"fmt"
	"math"
)

// pipeline - конвейер чисел.
// Читает числа из in, преобразует их в float64, возводит в куб и записывает в out.
func pipeline(in <-chan uint8, out chan<- float64) {
	go func() {
		defer close(out) // Закрываем выходной канал после завершения работы.
		for num := range in {
			floatNum := float64(num)
			cube := math.Pow(floatNum, 3)
			out <- cube
		}
	}()
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	// Запускаем конвейер.
	pipeline(in, out)

	// Запускаем горутину, которая отправляет числа в канал in.
	go func() {
		defer close(in) // Закрываем входной канал после отправки всех чисел.
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
	}()

	// Читаем результаты из канала out и выводим их на экран.
	for result := range out {
		fmt.Println(result)
	}
}
