package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func write(h int, wg *sync.WaitGroup, name string, size map[string]int) {
	defer wg.Done()
	mu.Lock()
	size[name] = h
	mu.Unlock()
	time.Sleep(1 * time.Second)
}

func main() {
	size := make(map[string]int)
	var wg sync.WaitGroup
	names := []string{"Ann", "Paul", "Andrew"}
	height := []int{157, 175, 190}

	for i := 0; i < len(names); i++ {
		wg.Add(1)
		go write(height[i], &wg, names[i], size)
	}

	wg.Wait()
	time.Sleep(1 * time.Second)

	fmt.Print(size)
}
