package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	//ioCopy()
	reader1()
	//goInAction1()
	//lyndaWrite()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func bufio() {

}

func ioCopy() {
	// Reading already created file and changing its content
	// O_APPEND mean the content which will be written will not
	// replace the previous content but will append it
	//file, err := os.Open("./programatically_created_file")    <- if file opened like this, file by default will opened as ReadOnly
	// refer - https://developpaper.com/solution-to-golang-bad-file-descriptor-problem/
	file, err := os.OpenFile("./programatically_created_file", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	checkErr(err)

	// Writing in to the file through file.Write method
	//content := "Hello this is a programatically created content in this file which appends previous content."
	//_, err = file.Write([]byte(content))
	//checkErr(err)

	// Writing though os.Copy()
	content := new(bytes.Buffer)
	content.WriteString("This content is writen by os.Copy method.")
	_, err = io.Copy(file, content)
	checkErr(err)
	
}

func lyndaWrite() {

	// Creating a file in the current directory by using os.Create method
	file, err := os.Create("./programatically_created_file")
	// Closing the file at the end hence used 'defer'
	defer file.Close()
	// Checking for errors
	checkErr(err)

	stringToWrite := "This string is written in this file programatically"
	// writing the 'stringToWrite' string to the file
	file.WriteString(stringToWrite)

	// writing to the file using ioutil package
	b := []byte(stringToWrite)
	ioutil.WriteFile("created_by_ioutil", b, 0777)

}

func goInAction1() {
	// creating a Buffer variable and will write a string to it
	// the Write method implements io.Writer
	var b bytes.Buffer
	b.Write([]byte("Hello this is learning IO in Golang"))

	// Concatinating string with a Buffer
	// passing address of Buffer for io.Writer
	fmt.Fprintf(&b, " and learning golang is awesome")

	// Writing to os.File. Stdin, Stdout, and Stderr are open Files
	// pointing to the standard input, standard output, and standard error file descriptors
	// passing the address of os.File, as WriteTo method takes *File as argument.
	b.WriteTo(os.Stdout)
	
}

func reader1() {
	// Using os package to open a file which will return  *os.File, don't forget to close the file - r.Close()
	r, _ := os.Open("file.txt")
	defer r.Close() // Closing the file
	//Reading through ioutil.ReadAll method
	b, _ := ioutil.ReadAll(r)
	fmt.Println("Through ioutil -", string(b))

	// Reading through io.Reader.Read() method
	buff := make([]byte, 256) // If we create a varialbe by make(), the Read() method works.
	//var b []byte    // if we define variable like this, it does not read.
	r.Read(buff)
	fmt.Println(string(buff))
	/*
		// ioutil.ReadFile("filename.txt") returns []byte .
		content, err := ioutil.ReadFile("file.txt")
		// checking if the ReadFile threw any error
		if err != nil {
			panic(err)
		}
		// Printing the string form of the content which is []byte
		fmt.Println(string(content))
	*/

	// Reading keeping the buffer size low
	content := new(bytes.Buffer)
	content.WriteString("This is going to be writen slowly because I have added a sleep statement, so the program would sleep for 1 second")
	buffer := make([]byte, 4)
	for {

		_, err := content.Read(buffer)
		if err == io.EOF {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf(string(buffer))
	}
}
