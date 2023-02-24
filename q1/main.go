package main

import (
	"fmt"
)

type Shape interface {
	area()float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func printArea(s Shape){
	fmt.Println("Area:", s.area())
}

func main() {
	r := rectangle{width: 7, height: 5}
	c := circle{radius: 10}
	printArea(r)
	printArea(c)
}
