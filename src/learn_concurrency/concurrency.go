package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func print(s string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	wg.Add(1)
	go print("anshul")
	wg.Add(1)
	go print("madhu")
	wg.Wait()
}
