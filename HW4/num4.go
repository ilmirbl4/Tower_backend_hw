package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, works <-chan string) {
	defer wg.Done()
	for work := range works {
		fmt.Printf("Результат работы воркера %d: %s", id+1, work)
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста укажите количество воркеров")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Введено недостаточное количество воркеров")
		return
	}

	works := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, works)
	}

	go func() {
		for {
			work := fmt.Sprintf("Задача %d", rand.Intn(322))
			works <- work
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sgn
		close(works)
	}()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
