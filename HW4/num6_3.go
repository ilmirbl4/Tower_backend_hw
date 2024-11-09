package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ch <-chan struct{}) {
	defer wg.Done() // уменьшаем счетчик
	for {
		select {
		case <-ch:
			fmt.Println("Goroutine stopped")
			return
		default:
			// Выполнение основной работы
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}) //создаем канал, который может передавать пустые структуры

	wg.Add(1) // увеличиваем счетчик горутин
	go worker(&wg, ch)

	// Даем горутине поработать 5 секунд
	time.Sleep(5 * time.Second)

	// Отправляем сигнал о завершении
	close(ch)

	// Ждем завершения горутины
	wg.Wait()
}
