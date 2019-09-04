package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	//maxp := flag.Int("-i", 6, "the max value")
	str := flag.String("j", "Hello", "-m")

	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(*str)
}
