package main

import "fmt"
import "time"

var start time.Time

func init() {
	start = time.Now()
}

func sum(x int) {
	/*
		sum := 0
		for i := 0; i < x; i++ {
			sum += i
		}
		fmt.Println("Sum of", x, "is", sum, "Time:", time.Since(start))
	*/
	fmt.Println(x, "time:", time.Since(start))
}
func scan() {
	fmt.Println("Scan function")
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(a)
}

func main() {
	go sum(100)
	//go scan()
	//fmt.Println("hello")
	go sum(200)
	go sum(300)
	go sum(400)
	time.Sleep(10 * time.Millisecond)
	//fmt.Scanln()
}
