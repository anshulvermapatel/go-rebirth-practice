package main

import "fmt"
import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

type Rect struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func (c Cube) Area() float64 {
	return 6 * c.side * c.side
}

func (c Cube) Perimeter() float64 {
	return 12 * c.side
}

func (r Rect) Area() float64 {
	return r.length * r.width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * (math.Pi) * c.radius
}

func main() {
	/*
		var s Shape = Circle{7}
		fmt.Println(s.Area())
	*/
	var o Object = Cube{3}
	fmt.Printf("o is of type %T and volume of the cube is %f\n", o, o.Volume())
	c := o.(Cube) // this is called type assertion, this is used when you have interface but you want to access the value inside the struct it refers to and also
	// it refers to and also it is used to convert the struct to any other type(interface) which it implements
	s, doesImplements := o.(Shape) // this is how we see that if Cube also implements Shape, doesImplements is a bool type, true if it implements else false.
	fmt.Printf("c is of type %T and its area is %f", c, c.Area())
	fmt.Printf("\nType of doesImplements is %T\n", doesImplements)
	if doesImplements {
		fmt.Println("Yes Cube also Implements Shape")
	} else {
		fmt.Println("No Cube does not Implements Shape")
	}
	fmt.Printf("s is of type %T and its area is %f\n", s, s.Area())
	fmt.Println(s.(type))
}
