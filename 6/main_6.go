package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// randomGenerator генерирует случайные числа и отправляет их в небуферизированный канал.
// Работает до тех пор, пока не будет отменен контекст.
func randomGenerator(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch) // Закрываем канал после завершения горутины.
		for {
			select {
			case <-ctx.Done():
				return // Выходим из горутины, если контекст отменен.
			case ch <- rand.Intn(100): // Отправляем случайное число в канал.
			}
		}
	}()

	return ch
}

func main() {

	// Создаем контекст с отменой.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Отменяем контекст после завершения main.

	// Получаем канал от генератора.
	randomCh := randomGenerator(ctx)

	// Читаем числа из канала в течение 5 секунд.
	timeout := time.After(5 * time.Second)
	for {
		select {
		case num := <-randomCh:
			fmt.Println("Получено случайное число:", num)
		case <-timeout:
			fmt.Println("Время вышло, завершаем работу.")
			return
		}
	}
}
