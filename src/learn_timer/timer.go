package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 10
	c := time.After(time.Second * time.Duration(timeout))
	for {
		if len(c) > 0 {
			fmt.Println("len(c)", len(c), "channel -", <-c)
			return
		}
		//fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
