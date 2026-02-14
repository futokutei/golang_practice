package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	SideBC float64
	SideAB float64
}

func (r Rectangle) Area() float64 {
	return r.SideAB * r.SideBC
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.SideAB + r.SideBC)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Triangle struct {
	SideAB float64
	SideBC float64
	SideAC float64
}

func (t Triangle) Perimeter() float64 {
	return t.SideAB + t.SideBC + t.SideAC
}

func (t Triangle) Area() float64 {
	p := (t.SideAB + t.SideBC + t.SideAC) / 2.0
	return math.Sqrt((p) * (p - t.SideAB) * (p - t.SideBC) * (p - t.SideAC))
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var result []int
	var b []int
	for i := 0; i < 10; i++ {
		b = append(b, int(rand.Int32N(100)))
	}
	for i := 0; i < 10; i++ {
		result = append(result, a[i]+b[i])
		fmt.Printf("Array 'а' element: %d, slice 'b' element: %d, slice 'result' element: %d\n", a[i], b[i], result[i])
	}

	r := Rectangle{SideAB: 10.0, SideBC: 20.0}
	c := Circle{Radius: 5.0}
	t := Triangle{SideAC: 10.0, SideBC: 10.0, SideAB: 10.0}
	fmt.Println(" ")
	fmt.Printf("Rectangle perimeter: %.2f, rectangle area: %.2f\n", r.Perimeter(), r.Area())
	fmt.Printf("Circle perimeter: %.2f, circle area: %.2f\n", c.Perimeter(), c.Area())
	fmt.Printf("Triangle perimeter: %.2f, triangle area: %.2f\n", t.Perimeter(), t.Area())

}
