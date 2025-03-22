package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := randomGenerator(ctx)

	// Проверяем, что из канала можно прочитать числа.
	select {
	case num := <-ch:
		t.Logf("Получено число: %d", num)
	case <-time.After(1 * time.Second):
		t.Error("Не удалось получить число из канала за 1 секунду")
	}

	// Проверяем, что канал закрывается при отмене контекста.
	cancel()
	select {
	case _, ok := <-ch:
		if ok {
			t.Error("Канал не был закрыт после отмены контекста")
		} else {
			t.Log("Канал успешно закрыт после отмены контекста")
		}
	case <-time.After(1 * time.Second):
		t.Error("Канал не закрылся за 1 секунду")
	}
}

func TestRandomGeneratorMultipleNumbers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := randomGenerator(ctx)

	// Проверяем, что из канала можно прочитать несколько чисел.
	for i := 0; i < 5; i++ {
		select {
		case num := <-ch:
			t.Logf("Получено число: %d", num)
		case <-time.After(1 * time.Second):
			t.Errorf("Не удалось получить число %d из канала за 1 секунду", i+1)
			return
		}
	}
}
