package main

import (
	"fmt"
	"sync"
)

func squares(x int, wg *sync.WaitGroup, res chan<- int) {
	defer wg.Done() // уменьшаем счетчик
	x = x * x
	res <- x // записываем квадрат числа в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup               // создаем счетчик горутин
	res := make(chan int, len(numbers)) // создаем канал

	for i := 0; i < len(numbers); i++ {
		wg.Add(1) // увеличиваем счетчик на 1
		go squares(numbers[i], &wg, res)
	}

	wg.Wait()  // ждем пока отработают все горутины, т.е. счетчик станет равным 0
	close(res) // закрываем канал

	sum := 0
	for square := range res {
		sum += square // суммируем квадраты чисел
	}

	fmt.Print(sum)
}
