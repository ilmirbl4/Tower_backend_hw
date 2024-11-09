package main

import (
	"fmt"
	"time"
)

func worker(ch chan struct{}) {
	for {
		select {
		case <-ch: // если подан сигнал о завершении
			fmt.Println("Горутина остановлена")
			close(ch) // закрываем канал
		default:
			fmt.Println("Я работаю...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan struct{}) // создаем канал, который может передавать пустые структуры

	go worker(ch)

	time.Sleep(5 * time.Second) // даем горутине поработать 5 секунд

	close(ch) // отправляем сигнал о завершении

	time.Sleep(1 * time.Second) // Ждем завершения горутины
}
