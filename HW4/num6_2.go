package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // если контекст отменен:
			fmt.Println("Горутина остановлена")
			return
		default:
			fmt.Println("Я работаю...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	//создаем новый контекст, который будет отменен после вызова cancel()
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)
	time.Sleep(5 * time.Second) // даем горутине поработать

	cancel() // вызываем cancel, чтобы отменить контекст
	time.Sleep(1 * time.Second)
}
