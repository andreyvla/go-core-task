package main

import (
	"fmt"
	"sync"
	"time"
)

// merge сливает N каналов в один.
func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Запускаем output горутину для каждого входного канала.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Запускаем горутину, которая закрывает выходной канал после завершения всех output горутин.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Создаем несколько каналов.
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	// Запускаем горутины, которые отправляют данные в каналы.
	go func() {
		defer close(c1)
		for i := 1; i <= 5; i++ {
			c1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer close(c2)
		for i := 10; i <= 15; i++ {
			c2 <- i
			time.Sleep(150 * time.Millisecond)
		}
	}()

	go func() {
		defer close(c3)
		for i := 20; i <= 23; i++ {
			c3 <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// Сливаем каналы в один.
	merged := merge(c1, c2, c3)

	// Читаем данные из объединенного канала и выводим их на экран.
	for n := range merged {
		fmt.Println(n)
	}
}
