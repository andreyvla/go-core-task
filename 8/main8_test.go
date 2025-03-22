package main

import (
	"sync"
	"testing"
	"time"
)

// TestCustomWaitGroup проверяет корректную работу CustomWaitGroup.
func TestCustomWaitGroup(t *testing.T) {
	var wg CustomWaitGroup
	wg.Add(3)

	var wgInner sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wgInner.Add(1)
		go func(id int) {
			defer wgInner.Done()
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
		}(i)
	}

	wgInner.Wait()
	wg.Wait()
}

// TestCustomWaitGroupAddDone проверяет корректную работу методов Add и Done.
func TestCustomWaitGroupAddDone(t *testing.T) {
	var wg CustomWaitGroup
	wg.Add(5)
	if wg.counter != 5 {
		t.Errorf("Ожидается, что счетчик будет равен 5, получено %d", wg.counter)
	}
	wg.Done()
	if wg.counter != 4 {
		t.Errorf("Ожидается, что счетчик будет равен 4, получено %d", wg.counter)
	}
}

// TestCustomWaitGroupWait проверяет корректную работу метода Wait.
func TestCustomWaitGroupWait(t *testing.T) {
	var wg CustomWaitGroup
	wg.Add(1)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Error("Wait вернулся до вызова Done")
	case <-time.After(100 * time.Millisecond):
		// Ожидаемое поведение
	}

	wg.Done()

	select {
	case <-done:
		// Ожидаемое поведение
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait не вернулся после вызова Done")
	}
}

// TestCustomWaitGroupNegativeCounter проверяет панику при отрицательном счетчике.
func TestCustomWaitGroupNegativeCounter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Код не вызвал панику")
		}
	}()

	var wg CustomWaitGroup
	wg.Add(-1)
}
