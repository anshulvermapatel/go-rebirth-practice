/*package main

import "fmt"

func main() {
	//	OUTPUT :-
		fmt.Println("This is " + "Anshul")
		fmt.Println("\nI am gonno add some numbers here 2 + 5 = ", 2+5)
		fmt.Println("\nThis is gonno be true ", true || false)
		fmt.Println("\nThis is gonno be false ", true && false)
		fmt.Printf("\nThis is gonno be %t ", false || false)
		fmt.Println("\nThis is gonno be true ", true || true)
		fmt.Println("\nThis is gonno be true ", true && true)
		fmt.Println("\nThis is gonno be false ", false && false)

*/
// Variables :-
/*
	var anshul string
	fmt.Println("anshul is - ", anshul)
	anshul = "A great man"
	fmt.Printf("Now anshul is %s\n", anshul)
	var a, b int = 50, 30
	fmt.Println("a, b are", a, "and", b, "respectively")
	c, d := 10, 20
	fmt.Println(c + d)
*/

// Constants :-
//const n = 400
//fmt.Println(n / 2)

// Looping
/*
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
*/
// while like for loop
/*
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
*/
// iterate over only even numbers :-
/*
	for i := 1; i < 30; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
*/
// infinite loop
/*
	for {
		fmt.Println("this is going to run forever")
	}
*/

// if/else with Scan and Scanf
/*
	var a int
	var b bool
	fmt.Println("Do you want to perform if else operation?")
	fmt.Scanf("%t", &b)
	if b == false {
		return
	}
	fmt.Println("Give a for the Operation")
	fmt.Scan(&a)
	//a = 40
	if a > 50 {
		fmt.Println("a is greater than 50")
	} else if a < 30 {
		fmt.Println("a is less then 30")
	} else {
		fmt.Println("a is greater than 30 and less than 50")
	}
	fmt.Println("b is", b)
*/
// Switch
/*
	switch {
	case 5 > 6: // (5 > 6) this has to be boolean when you dont give value for switch to compare with.
		fmt.Println("hello")
	default:
		fmt.Println("default")
	}
*/

// Arrays
/*
	var a = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(a)
	}
	for _, b := range a {
		fmt.Println(b)
	}
	weekdays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i, weekday := range weekdays {
		fmt.Println(i, "Day is", weekday)
	}
*/
// Slices
/*
			a := make([]int, 5)
			a = []int{1, 2, 3, 4, 5}
			for _, b := range a {
				fmt.Println(b)
			}
			b := a[0:3]
			fmt.Println(b)
			a[0] = 5
			fmt.Println(b)
			a = append(a, 10)
			a = append(a, 15)
			fmt.Println("Length of a is ", len(a))
			fmt.Println(a)

		a := [...]int{1, 2, 3, 4, 5} // this is for array
		b := []int{1, 2, 3, 4, 5}    // this is for slice
		fmt.Println(reflect.TypeOf(a))
		//a = append(a, 7)
		fmt.Println(a)
		fmt.Println(b)

	a := make([]int, 5, 10)
	a = []int{0, 1, 2, 3, 4}
	a = append(a, 5)
	fmt.Println("Length of a is ", len(a))
	a = append(a, 6)
	a = append(a, 6)
	a = append(a, 6)
	a = append(a, 6)
	fmt.Println("Length of a is ", len(a))
	fmt.Println("Capacity of a is ", cap(a))
}
*/
