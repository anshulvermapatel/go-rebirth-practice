package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/*
		a := "Madhulika Choubey"
		fmt.Println("Before: -", a)
		// Print the string in upper case
		fmt.Println(strings.ToUpper(a))
		// Print the string in lower case
		fmt.Println(strings.ToLower(a))

		fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, a))

		// Prints if the string has an string in its prefix - strings.HasPrefix(string, prefix_string)
		fmt.Println(strings.HasPrefix(a, "Madhulika"))
		// Prints if the string has an string in its suffix - strings.HasPrefix(string, suffix_string)
		fmt.Println(strings.HasSuffix(a, "Choubey"))

		// returns boolean if a string contains a string given for search, it is case sensitive
		a = "Hello this is Anshul Verma Verma"
		fmt.Println(strings.Contains(a, "Hell"))
		fmt.Println(strings.Contains(a, "Is"))

		// return how many time a string apears in the string
		fmt.Println(strings.Count(a, "Verma"))
	*/

	/*
	  **********************************
	  Functions for String Manupulation
	  **********************************
	*/

	/*
		// Used to join the strings in a slice with being saperated by a character given as string in the next argument
		fmt.Println(strings.Join([]string{"Hello", "I am", "Anshul"}, ":"))

		//lexicographically outputs -1, 0 , 1
		fmt.Println(strings.Compare("Anshul", "Khushal"))  // -1
		fmt.Println(strings.Compare("Khushal", "Anshul"))  // 0
		fmt.Println(strings.Compare("Khushal", "Khushal")) // 1

		//Prints true if any string found whether any unicode code points in chars
		fmt.Println(strings.ContainsAny(a, "s"))     // true
		fmt.Println(strings.ContainsAny(a, "z & z")) // true
		fmt.Println(strings.ContainsAny(a, "s & z")) // true
		fmt.Println(strings.ContainsAny(a, "z"))     // false
	*/

	/*
		start := time.Now()
		var b strings.Builder
		for i := 3; i >= 1; i-- {
			fmt.Fprintf(&b, "%d...", i)
		}
		b.WriteString("ignition")
		fmt.Println(b.String())
		//fmt.Println("ignition")
		fmt.Println(time.Since(start))
	*/

	// Trims
	// returns a slice of string made by cutting a string using a
	fmt.Println(strings.Trim("???/// Hello Girl ////???", "?/"))        // [ Hello Girl ]
	fmt.Println(strings.Trim("???///Hello Girl ////???", "/?"))         // [Hello Girl]
	fmt.Println(strings.Trim("??? /// Hello Girl ////???", "?/"))       // [ /// Hello Girl ] The cutset has to be together, in not only work if the starting of the string is in the cutset
	fmt.Println(strings.Trim("/// ??? Hello ///?? Girl ////???", "/?")) // [ /// Hello Girl ]

	// returns string after removing all the white spaces around the string passed in the argument.
	fmt.Printf(strings.TrimSpace("  Hello, I am Anshul    ")) // Hello, I am Anshul

	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}
