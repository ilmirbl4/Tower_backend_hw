package main

import (
	"fmt"
	"sync"
)

func squares(x int, wg *sync.WaitGroup, res chan<- int) {
	defer wg.Done() // после выполнения уменьшаем счетчик горутин
	x = x * x
	res <- x // записываем квадрат числа в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup               // создаем вэйтгруппу
	res := make(chan int, len(numbers)) // создаем канал, в который горутины будут передавать результаты

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)                        // увеличиваем счетчик горутин
		go squares(numbers[i], &wg, res) // запускаем горутину
	}

	wg.Wait()  // ждем пока все горутины отработают
	close(res) // закрываем канал

	for square := range res {
		fmt.Print(square, " ") // вывод результатов
	}
}
