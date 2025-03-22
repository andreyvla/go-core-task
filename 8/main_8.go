package main

import (
	"fmt"
	"sync"
	"time"
)

// CustomWaitGroup - кастомная реализация WaitGroup на семафоре.
type CustomWaitGroup struct {
	counter   int
	semaphore chan struct{}
	mutex     sync.Mutex
}

// Add увеличивает счетчик горутин.
func (wg *CustomWaitGroup) Add(delta int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.counter += delta
	if wg.counter < 0 {
		panic("Отрицательный WaitGroup счетчик")
	}

	if wg.counter == 0 && wg.semaphore != nil {
		close(wg.semaphore)
		wg.semaphore = nil
	}
}

// Done уменьшает счетчик горутин на 1.
func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

// Wait блокирует выполнение, пока счетчик не станет равен 0.
func (wg *CustomWaitGroup) Wait() {
	wg.mutex.Lock()
	if wg.counter == 0 {
		wg.mutex.Unlock()
		return
	}
	if wg.semaphore == nil {
		wg.semaphore = make(chan struct{})
	}
	wg.mutex.Unlock()
	<-wg.semaphore
}

func main() {
	var wg CustomWaitGroup
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d начала работу\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Горутина %d завершила работу\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}
