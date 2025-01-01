package main

import (
	"fmt"
	"time"
)

func writing(ch chan<- int, workTime time.Duration) {
	startTime := time.Now()
	count := 0

	for time.Since(startTime) < workTime {
		ch <- count
		time.Sleep(1 * time.Second)
		count += 1
	}

	close(ch)
}

func reading(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	var workTime int
	fmt.Print("Пожалуйста, введите сколько сколько секунд должена работать программа: ")
	fmt.Scanf("%d", &workTime)

	ch := make(chan int)
	go writing(ch, time.Duration(workTime)*time.Second)
	reading(ch)

	fmt.Println("Программа завершена")
}
