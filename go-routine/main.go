package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d start \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d end \n", id)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

	fmt.Println("total time: ", time.Since(start))
}
