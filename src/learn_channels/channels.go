package main

import (
	"fmt"
	"sync"
)

// taking out the value from channel work in LIFO (Last In First Out)

var wg sync.WaitGroup

func double(num int, channel chan int) {
	defer wg.Done()
	fmt.Println(num*2, "is going into the channel")
	channel <- num * 2
}

func main() {
	channel := make(chan int, 4) // I need to use buffer(the second argument of the make function) here, I don't why it has to be used, without this it was throwing an error.
	wg.Add(1)
	go double(5, channel)
	wg.Add(1)
	go double(10, channel)
	wg.Add(1)
	go double(20, channel)
	wg.Add(1)
	go double(30, channel)

	wg.Wait()
	close(channel)   // if you remove this you will get an error - "fatal error: all goroutines are asleep - deadlock!", this is required only when iterating over channel.
	//v1 := <-channel
	//v2 := <-channel
	for i := range channel {
		fmt.Println(i)
	}

	/*
		channel := make(chan int)
		go double(5, channel)
		go double(10, channel)
		go double(20, channel)
		go double(30, channel)
		//v1 := <-channel
		//v2 := <-channel
		for i := range channel {
			fmt.Println(i)
		}
		//fmt.Println("v1:", v1, "| v2:", v2)
	*/
}
