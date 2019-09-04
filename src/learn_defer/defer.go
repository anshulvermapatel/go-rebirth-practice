package main

import "fmt"

func hello(x []int, index int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("There is a panic: Array out of bound")
		}
	}()

	if len(x) > index {
		return x[index]
	} else {
		panic("Array out of bound")
	}
	return 0
}

/*
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("There is a panic: Array out of bound")
	}
}
*/
func main() {
	x := []int{1, 2, 3}
	//defer recovery()
	fmt.Println(hello(x, 3))
}
