/*
package main

import "fmt"

type Salaried interface {
	getSalary() int
}

type Salary struct {
	base      int
	insurance int
	allowance int
}

func (s Salary) getSalary() int {
	return s.base + s.allowance + s.insurance
}

type Employee struct {
	firstName, lastName string
	Salaried
}

type Person struct {    // the field type cannot be repeat, e.g there cannot be 2 string fields or any other field, one type can only be once in Anonymous field type struct.
	string
	int
	bool
}

func main() {

	//

	// Anonymous fields  see the struct type Person
	/*
	p := Person{"Anshul", 24, true}
	fmt.Println(p)
*/
// pointer to a struct
/*
	s := &Salary{20, 25, 30}
	fmt.Println((*s).base) // *(s.base) does not work but s.base will work because it get dereferenced automatically
*/
// Anonymous struct
/*
	monica := struct {
		name       string
		age        int
		proffesion string
	}{
		name:       "Anajali",
		age:        23,
		proffesion: "Doctor",
	}
	fmt.Println(monica)
*/
/*
	ross := Employee{
		firstName: "Anshul",
		lastName:  "Verma",
		//	Salaried:  Salary{base: 58000, insurance: 500000, allowance: 12000},  // if this is commented out so ross.Salaried will be <nil> and if called getSalary() go will panic.
	}
	//fmt.Println(ross.getSalary())
	fmt.Println("doing nothing", ross.Salaried)
*/
/*
	b := &Salary{base: 58000, insurance: 500000, allowance: 12000}
	var s Salaried
	s = b
	//s.base = 5
	//s.insurance = 10
	//s.allowance = 10
	b.base = 5
	b.allowance = 5
	b.insurance = 5
	fmt.Println(s.getSalary())
*/
//}
