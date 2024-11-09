package main

import (
	"fmt"
	"sync"
	"time"
)

var stopFlag bool
var mu sync.Mutex // создаем новый мьютекс

func worker() {
	for {
		mu.Lock() // закрывем доступ
		if stopFlag {
			mu.Unlock()
			fmt.Println("Горутина остановлена")
			return
		}
		mu.Unlock() // открываем доступ

		// Выполнение основной работы
		fmt.Println("Я работаю...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go worker()

	// Даем горутине поработать 5 секунд
	time.Sleep(5 * time.Second)

	// Устанавливаем флаг завершения
	mu.Lock()
	stopFlag = true
	mu.Unlock()

	time.Sleep(1 * time.Second)
}
